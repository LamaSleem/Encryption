
build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/room-one \
	  proto/encryption.proto
docker:
	docker build -t room-one .
kubectl:
    kubectl apply -f encryption-service.yaml