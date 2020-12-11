FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression. Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .

FROM golang:1.15 AS builder

# Services:
RUN apt-get update && apt-get install -y openssh-client

# Preparing project env:
RUN mkdir -m 0777 -p /go/src/github.com/spyzhov/grpc_mock

WORKDIR /go/src/github.com/spyzhov/grpc_mock/mock_gen

ENV GOOS=linux
ENV GOARCH=amd64

COPY ./mock_gen/go.mod ./mock_gen/go.sum ./
RUN go mod download

COPY ./mock_gen .

RUN go mod vendor -v
RUN go build -mod=vendor -o "/go/bin/mock_gen" "."

FROM golang:1.15

# environment
ENV MANAGE_PORT=8000
ENV GRPC_PORT=8090
ENV PROTO_PATH=/proto
ENV PROTO_PATHS="/proto"

# configurations
EXPOSE 8000
EXPOSE 8090
VOLUME /proto

# Services:
RUN apt-get update && apt-get install -y openssh-client protobuf-compiler
RUN GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@v1.4.3
#RUN GO111MODULE=on \
#    go get google.golang.org/protobuf/cmd/protoc-gen-go \
#             google.golang.org/grpc/cmd/protoc-gen-go-grpc

# the timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# the mock_gen program:
COPY --from=builder /go/bin/mock_gen /usr/bin/mock_gen

WORKDIR /go/src/github.com/spyzhov/grpc_mock/
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

ENTRYPOINT ["bash",  "docker-entrypoint.sh"]
CMD ["grpc_mock"]
