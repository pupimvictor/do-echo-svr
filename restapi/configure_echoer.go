// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/pupimvictor/do-echo-svr"
	"net/http"
	"os"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/pupimvictor/do-echo-svr/restapi/operations"
	"github.com/pupimvictor/do-echo-svr/restapi/operations/echo"

	models "github.com/pupimvictor/do-echo-svr/models"
)

//go:generate swagger generate server --target ../../do-echo-svr --name Echoer --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.EchoerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

//EchoHandlerFn is the handler function for the /echo POST request. (See swagger.yml for spec details)
//It calls echoer.Echo to perform the echo logic
//It will return a http 400 code in the case of a request without a body
func EchoHandlerFn(params echo.EchoParams, principal *models.Principal) middleware.Responder {
	if params.Body != nil {
		resp := echoer.Echo(params.Body)
		return echo.NewEchoOK().WithPayload(resp)
	}
	return echo.NewEchoDefault(http.StatusBadRequest)
}

func configureAPI(api *operations.EchoerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-Token" header is set
	api.TokenHeaderAuth = func(token string) (*models.Principal, error) {
		if token == os.Getenv("API_TOKEN") {
			prin := models.Principal(token)
			return &prin, nil
		}
		return nil, errors.New(401, "incorrect api key")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.EchoEchoHandler = echo.EchoHandlerFunc(EchoHandlerFn)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
