package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"net/http"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	"github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/protobuf/types"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/nzoschke/gomesh/internal/metadata"
)

// Interface assertion
var _ auth.AuthorizationServer = (*Server)(nil)

// Server implements the auth/v2alpha interface
type Server struct {
	Hydra     *hydra.CodeGenSDK
	Transport *Transport
}

// Transport adds host and trace headers to requests
type Transport struct {
	Authority string
	TraceID   string
	Transport *http.Transport
}

// RoundTrip implements the RoundTripper interface
// It rewrites the host for envoy forwarding
// Andy hackily sets a trace id
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = t.Authority

	req.Header.Add("uber-trace-id", t.TraceID)
	resp, err := t.Transport.RoundTrip(req)
	t.TraceID = ""

	return resp, err
}

// Check does an auth header check and returns a 200 on auth, non-200 otherwise
func (s *Server) Check(ctx context.Context, in *auth.CheckRequest) (*auth.CheckResponse, error) {
	h := in.Attributes.Request.Http.Headers["authorization"]
	if h == "" {
		return responseDenied("No auth header present", int(envoy_type.StatusCode_Unauthorized))
	}

	auth := strings.SplitN(h, " ", 2)

	switch auth[0] {
	case "Basic":
		return s.basicCheck(auth[1])
	case "Bearer":
		return s.bearerCheck(ctx, auth[1])
	default:
		return responseDenied("Invalid auth header", int(envoy_type.StatusCode_Unauthorized))
	}
}

func (s *Server) basicCheck(token string) (*auth.CheckResponse, error) {
	payload, _ := base64.StdEncoding.DecodeString(token)
	parts := strings.SplitN(string(payload), ":", 2)

	if len(parts) != 2 || !basicValidate(parts[0], parts[1]) {
		return responseDenied("Invalid basic credentials", int(envoy_type.StatusCode_Unauthorized))
	}

	return responseOk(fmt.Sprintf("users/%s", parts[0]))
}

// TODO: implement real password checking
func basicValidate(username, password string) bool {
	return true
}

func (s *Server) bearerCheck(ctx context.Context, token string) (*auth.CheckResponse, error) {
	tid, _ := metadata.Get(ctx, "uber-trace-id")
	s.Transport.TraceID = tid

	i, _, err := s.Hydra.IntrospectOAuth2Token(token, "")
	if err != nil {
		return responseDenied(err.Error(), int(envoy_type.StatusCode_Unauthorized))
	}

	if !i.Active {
		return responseDenied("Inactive bearer token", int(envoy_type.StatusCode_Unauthorized))
	}

	return responseOk(fmt.Sprintf("clients/%s", i.ClientId))
}

func responseOk(subject string) (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					&core.HeaderValueOption{
						Append: &types.BoolValue{
							Value: false,
						},
						Header: &core.HeaderValue{
							Key:   "x-subject-id",
							Value: subject,
						},
					},
				},
			},
		},
		Status: &rpc.Status{
			Code: int32(rpc.OK),
		},
	}, nil
}

func responseDenied(body string, status int) (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Body: body,
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
			},
		},
		Status: &rpc.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
	}, nil
}