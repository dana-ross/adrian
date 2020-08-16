FROM golang

EXPOSE 80/tcp 443/tcp

RUN mkdir -p /go/src/github.com/daveross/adrian

ADD . /go/src/github.com/daveross/adrian

RUN cd /go/src/github.com/daveross/adrian && go get -v ./... && go build .

FROM alpine:latest

COPY --from=0 /go/src/github.com/daveross/adrian/adrian .

RUN apk add libc6-compat

RUN mkdir -p /fonts

ADD adrian.yaml.docker.example /etc/adrian.yaml

CMD ./adrian --config /etc/adrian.yaml