FROM golang:alpine as builder

WORKDIR /go/src/github.com/ebedevelopment/next-gen-tms/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="ebe@ebe.com.eg"

WORKDIR /go/src/github.com/ebedevelopment/next-gen-tms/server

COPY --from=0 /go/src/github.com/ebedevelopment/next-gen-tms/server/server ./
COPY --from=0 /go/src/github.com/ebedevelopment/next-gen-tms/server/resource ./resource/
COPY --from=0 /go/src/github.com/ebedevelopment/next-gen-tms/server/config.docker.yaml ./

EXPOSE 8889
EXPOSE 50051
EXPOSE 9100

ENTRYPOINT ./server -c config.docker.yaml
