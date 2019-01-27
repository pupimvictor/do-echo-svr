// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/pupimvictor/do-echo-svr/models"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/pupimvictor/do-echo-svr"
	"github.com/pupimvictor/do-echo-svr/restapi/operations"
	"github.com/pupimvictor/do-echo-svr/restapi/operations/echo"
)

//go:generate swagger generate server --target ../../do-echo-svr --name Echoer --spec ../swagger.yml

func configureFlags(api *operations.EchoerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func EchoHandlerFn(params echo.EchoParams) middleware.Responder {
	if params.Body != nil {
		resp := echoer.Echo(params.Body)
		return echo.NewEchoOK().WithPayload(resp)
	}
	errMsg := "please, say something!"
	return echo.NewEchoDefault(http.StatusBadRequest).WithPayload(&models.Error{Message: &errMsg})
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
