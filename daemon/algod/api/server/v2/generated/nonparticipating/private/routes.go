// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
