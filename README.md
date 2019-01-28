# Echo Server

This is a service that echoes back to a client when it is yelled at!

#### OpenAPI
The service (and [client](https://github.com/pupimvictor/do-echo-cli)) were created using [go-swagger](https://github.com/go-swagger/go-swagger/), a golang implementation of Swagger 2.0 (aka OpenAPI 2.0), to create a contract between the client and the server.

#### Structure
The go-swagger code generator provides a skeleton of the api based on the swagger.yml spec. I've place the heavy business logic of the Echo app at the root of the repository, but that might not be the best place for it.

The file restapi/operations/configure_echoer.go contains the handler functions and it's the place to hook up the generated code and the app business logic.

#### Authentication
You can use a x-api-token to authorize a request for an echo. make sure you pass it when deploying: `make deploy ENV=<path to a file with API_TOKEN=<token>>`

#### Usage
This app contains a Make file to help on the development and deployment of the service. It is setup to use Docker and to connect to a DigitalOcean Droplet using [Docker Machine](https://docs.docker.com/machine/)

To deploy the app, run `make deploy`.

