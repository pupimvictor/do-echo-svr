FROM alpine
RUN apk add --no-cache ca-certificates

ADD ./cmd/echoer-server/echoer-server /bin/echoer-server

ENTRYPOINT /bin/echoer-server --port 8000 --host="0.0.0.0"




