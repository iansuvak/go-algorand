// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
