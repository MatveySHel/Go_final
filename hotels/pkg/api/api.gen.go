//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
)

// Error Сообщение об ошибках
type Error struct {
	Message *string `json:"message,omitempty"`
}

// Hotel Hotel
type Hotel struct {
	// City Город
	City *string `json:"city,omitempty"`

	// Id ID
	Id *int `json:"id,omitempty"`

	// Name Название отеля
	Name *string `json:"name,omitempty"`

	// Price Цена
	Price *int `json:"price,omitempty"`
}

// CreateNewHotelJSONBody defines parameters for CreateNewHotel.
type CreateNewHotelJSONBody struct {
	// City Город
	City *string `json:"city,omitempty"`

	// Name Название отеля
	Name *string `json:"name,omitempty"`

	// Price Цена
	Price *int `json:"price,omitempty"`
}

// EditHotelJSONBody defines parameters for EditHotel.
type EditHotelJSONBody struct {
	// City Город
	City *string `json:"city,omitempty"`

	// Name Название отеля
	Name *string `json:"name,omitempty"`

	// Price Цена
	Price *int `json:"price,omitempty"`
}

// CreateNewHotelJSONRequestBody defines body for CreateNewHotel for application/json ContentType.
type CreateNewHotelJSONRequestBody CreateNewHotelJSONBody

// EditHotelJSONRequestBody defines body for EditHotel for application/json ContentType.
type EditHotelJSONRequestBody EditHotelJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получить массив с информацией об отеллях
	// (GET /hotels)
	GetHotelsList(w http.ResponseWriter, r *http.Request)
	// Создать новый отель
	// (POST /hotels)
	CreateNewHotel(w http.ResponseWriter, r *http.Request)
	// Изменение информации об отеле
	// (PUT /hotels)
	EditHotel(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetHotelsList operation middleware
func (siw *ServerInterfaceWrapper) GetHotelsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHotelsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateNewHotel operation middleware
func (siw *ServerInterfaceWrapper) CreateNewHotel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateNewHotel(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// EditHotel operation middleware
func (siw *ServerInterfaceWrapper) EditHotel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.EditHotel(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/hotels", wrapper.GetHotelsList)
	m.HandleFunc("POST "+options.BaseURL+"/hotels", wrapper.CreateNewHotel)
	m.HandleFunc("PUT "+options.BaseURL+"/hotels", wrapper.EditHotel)

	return m
}
