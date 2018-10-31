CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

configs/sidecar.yaml: cmd/envoy-cfg/main.go
	go run $< /tmp/sidecar.yaml
	docker run \
		-v/tmp/sidecar.yaml:/tmp/sidecar.yaml \
		envoyproxy/envoy:latest \
		envoy -c /tmp/sidecar.yaml --mode validate
	# mv /tmp/sidecar.yaml configs/sidecar.yaml

compose-api: bins
	docker-compose -f config/docker/compose-api.yaml -p gomesh build
	docker-compose -f config/docker/compose-api.yaml -p gomesh up

compose-mesh: bins
	docker-compose -f config/docker/compose-mesh.yaml -p gomesh build
	docker-compose -f config/docker/compose-mesh.yaml -p gomesh up

compose-proxy: bins
	docker-compose -f config/docker/compose-proxy.yaml -p gomesh build
	docker-compose -f config/docker/compose-proxy.yaml -p gomesh up --abort-on-container-exit

generate:
	docker build -f config/docker/Dockerfile-prototool -t prototool . 2>/dev/null
	docker run -v $(PWD):/in prototool /bin/prototool.sh

setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: vendor
vendor:
	git remote add -f -t master --no-tags protoc-gen-validate https://github.com/lyft/protoc-gen-validate.git || true
	git remote add -f -t master --no-tags grpc-gateway        https://github.com/grpc-ecosystem/grpc-gateway  || true
	git rm -rf vendor/
	git read-tree --prefix=vendor/github.com/lyft/protoc-gen-validate/validate/ -u protoc-gen-validate/master:validate
	git read-tree --prefix=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party -u grpc-gateway/master:third_party
