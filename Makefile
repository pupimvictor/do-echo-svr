.PHONY: build serve clean pack deploy ship

TAG?=v0.0.1

clean:
	cd ./cmd/echoer-server && rm ./echoer-server

build:
	cd cmd/echoer-server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o echoer-server .

pack:
	docker build -t pupimvictor/echoer-server .
	docker tag pupimvictor/echoer-server:latest pupimvictor/echoer-server:$(TAG)

upload:
	docker push pupimvictor/echoer-server:$(TAG)

ship: build pack upload

serve:
	docker run -p 8000:8000 pupimvictor/echoer-server:$(TAG)

deploy: build pack upload serve

droplet-new:
	docker-machine create --driver digitalocean --digitalocean-access-token $(do-token) echoer

droplet-ssh:
	eval $(docker-machine env echoer)
	docker-machine ssh echoer