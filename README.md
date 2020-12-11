# grpc_mock

Can be used for *.proto files with:

```protobuf
option go_package = "protob";
```

Usage:

```
docker run -it -v $(pwd)/helloworld/proto:/proto -p 8000:8000 -p 8090:8090 spyzhov/grpc_mock
```

For additional packages:

```
docker run -it -v $(pwd)/example/with_types:/proto -e PROTO_PATH=/proto/proto -e PROTO_PATHS="/proto/proto /proto/types" -p 8000:8000 -p 8090:8090 grpc_mock
```

Request:

```
grpcurl -plaintext -d '{"name": "word"}' localhost:8090 helloworld.Greeter/SayHello
```

Setup response:

```
curl -X POST -d '{"request":{"name": "word"}, "response": {"name": "Hello World!"}}' \
    localhost:8000/helloworld.Greeter/SayHello
```

Clear response:

```
curl -X DELETE localhost:8000/
```