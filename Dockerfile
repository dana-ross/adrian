FROM golang:1.15

EXPOSE 80/tcp 443/tcp

RUN mkdir -p /go/src/github.com/dana-ross/adrian

ADD . /go/src/github.com/dana-ross/adrian

RUN cd /go/src/github.com/dana-ross/adrian && go get -v ./... && go build .

FROM alpine:3

COPY --from=0 /go/src/github.com/dana-ross/adrian/adrian .

RUN apk add --no-cache libc6-compat

RUN mkdir -p /fonts

ADD adrian.yaml.docker.example /etc/adrian.yaml

CMD ./adrian --config /etc/adrian.yaml