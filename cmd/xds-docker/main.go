package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	xds "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port int `conf:"p" help:"Port to listen"`
}

// Server implements the Envoy xDS EndpointDiscovery service
type Server struct {
	lastVersion int
	lastNonce   string
}

func main() {
	config := config{
		Port: 8000,
	}
	conf.Load(&config)

	cli, err := client.NewEnvClient() // unix:///var/run/docker.sock by default
	if err != nil {
		panic(err)
	}
	log.Printf("CLIENT: %+v\n", cli)

	filter := filters.NewArgs()
	filter.Add("type", "container")
	filter.Add("event", "start")
	filter.Add("event", "die")

	since := time.Now().Add(-time.Minute)

	events, errs := cli.Events(context.Background(), types.EventsOptions{
		Since:   since.Format(time.RFC3339Nano),
		Filters: filter,
	})
	for {
		select {
		case err := <-errs:
			log.Fatal(err)
		case event := <-events:
			log.Printf("EVENT: %+v\n", event)
		}
	}
}

func serve(config config) error {
	s := grpc.NewServer()
	xds.RegisterEndpointDiscoveryServiceServer(s, &Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

// FetchEndpoints implements the Envoy xDS gRPC REST endpoint
func (s *Server) FetchEndpoints(ctx context.Context, request *xds.DiscoveryRequest) (*xds.DiscoveryResponse, error) {
	return nil, fmt.Errorf("unimplemented")
}

// StreamEndpoints implements the Envoy xDS gRPC streaming endpoint
func (s *Server) StreamEndpoints(server xds.EndpointDiscoveryService_StreamEndpointsServer) error {
	for {
		r, err := server.Recv()
		if err != nil {
			fmt.Printf("ERR: %+v\n", err)
			return err
		}
		fmt.Printf("R: %+v\n", r)

		// build response for new resources
		clas := []xds.ClusterLoadAssignment{}
		var as []ptypes.Any
		for _, cla := range clas {
			a, err := ptypes.MarshalAny(&cla)
			if err != nil {
				return err
			}
			as = append(as, *a)
		}

		// bump Nonce and Version
		s.lastNonce = string(time.Now().Nanosecond())
		s.lastVersion = s.lastVersion + 1

		server.Send(&xds.DiscoveryResponse{
			Resources:   as,
			Nonce:       s.lastNonce,
			VersionInfo: strconv.Itoa(s.lastVersion),
		})

		// Loop here reading docker events channel.
		// If anything has changed send the results to Envoy, otherwise do nothing and wait for the next batch of results.
		// for {
		// 	select {
		// 		case r
		// 	}
		// }
	}
}
