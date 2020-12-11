# grpc_mock

Usage:

```
docker run -it -v $(pwd)/helloworld/proto:/proto -p 8000:8000 -p 8090:8090 spyzhov/grpc_mock
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