// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package generated

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for LibraryBookResponseCondition.
const (
	LibraryBookResponseConditionBAD       LibraryBookResponseCondition = "BAD"
	LibraryBookResponseConditionEXCELLENT LibraryBookResponseCondition = "EXCELLENT"
	LibraryBookResponseConditionGOOD      LibraryBookResponseCondition = "GOOD"
)

// Defines values for ReturnBookRequestCondition.
const (
	ReturnBookRequestConditionBAD       ReturnBookRequestCondition = "BAD"
	ReturnBookRequestConditionEXCELLENT ReturnBookRequestCondition = "EXCELLENT"
	ReturnBookRequestConditionGOOD      ReturnBookRequestCondition = "GOOD"
)

// BookInfo defines model for BookInfo.
type BookInfo struct {
	// Author Автор
	Author string `json:"author"`

	// BookUid UUID книги
	BookUid openapi_types.UUID `json:"bookUid"`

	// Genre Жанр
	Genre string `json:"genre"`

	// Name Название книги
	Name string `json:"name"`
}

// ErrorDescription defines model for ErrorDescription.
type ErrorDescription struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

// LibraryBookPaginationResponse defines model for LibraryBookPaginationResponse.
type LibraryBookPaginationResponse struct {
	Items []LibraryBookResponse `json:"items"`

	// Page Номер страницы
	Page *int `json:"page,omitempty"`

	// PageSize Количество элементов на странице
	PageSize *int `json:"pageSize,omitempty"`

	// TotalElements Общее количество элементов
	TotalElements int `json:"totalElements"`
}

// LibraryBookResponse defines model for LibraryBookResponse.
type LibraryBookResponse struct {
	// Author Автор
	Author string `json:"author"`

	// AvailableCount Количество книг, доступных для аренды в библиотеке
	AvailableCount int `json:"availableCount"`

	// BookUid UUID книги
	BookUid openapi_types.UUID `json:"bookUid"`

	// Condition Состояние книги
	Condition LibraryBookResponseCondition `json:"condition"`

	// Genre Жанр
	Genre string `json:"genre"`

	// Name Название книги
	Name string `json:"name"`
}

// LibraryBookResponseCondition Состояние книги
type LibraryBookResponseCondition string

// LibraryPaginationResponse defines model for LibraryPaginationResponse.
type LibraryPaginationResponse struct {
	Items []LibraryResponse `json:"items"`

	// Page Номер страницы
	Page *int `json:"page,omitempty"`

	// PageSize Количество элементов на странице
	PageSize *int `json:"pageSize,omitempty"`

	// TotalElements Общее количество элементов
	TotalElements int `json:"totalElements"`
}

// LibraryResponse defines model for LibraryResponse.
type LibraryResponse struct {
	// Address Адрес библиотеки
	Address string `json:"address"`

	// City Город, в котором находится библиотека
	City string `json:"city"`

	// LibraryUid UUID библиотеки
	LibraryUid openapi_types.UUID `json:"libraryUid"`

	// Name Название библиотеки
	Name string `json:"name"`
}

// ReturnBookRequest defines model for ReturnBookRequest.
type ReturnBookRequest struct {
	// Condition Состояние книги
	Condition ReturnBookRequestCondition `json:"condition"`
}

// ReturnBookRequestCondition Состояние книги
type ReturnBookRequestCondition string

// ValidationErrorResponse defines model for ValidationErrorResponse.
type ValidationErrorResponse struct {
	// Errors Массив полей с описанием ошибки
	Errors []ErrorDescription `json:"errors"`

	// Message Информация об ошибке
	Message string `json:"message"`
}

// ViolationStatus defines model for ViolationStatus.
type ViolationStatus struct {
	// Violation Нарушены ли правила возврата книг
	Violation bool `json:"violation"`
}

// ListLibrariesParams defines parameters for ListLibraries.
type ListLibrariesParams struct {
	Page *int `form:"page,omitempty" json:"page,omitempty"`
	Size *int `form:"size,omitempty" json:"size,omitempty"`

	// City Город
	City string `form:"city" json:"city"`
}

// ListBooksParams defines parameters for ListBooks.
type ListBooksParams struct {
	Page    *int  `form:"page,omitempty" json:"page,omitempty"`
	Size    *int  `form:"size,omitempty" json:"size,omitempty"`
	ShowAll *bool `form:"showAll,omitempty" json:"showAll,omitempty"`
}

// ReturnBookJSONRequestBody defines body for ReturnBook for application/json ContentType.
type ReturnBookJSONRequestBody = ReturnBookRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получить информацию о книге
	// (GET /api/v1/books/{bookUid})
	GetBook(ctx echo.Context, bookUid openapi_types.UUID) error
	// Получить список библиотек в городе
	// (GET /api/v1/libraries)
	ListLibraries(ctx echo.Context, params ListLibrariesParams) error
	// Получить информацию о библиотеке
	// (GET /api/v1/libraries/{libraryUid})
	GetLibrary(ctx echo.Context, libraryUid openapi_types.UUID) error
	// Получить список книг в выбранной библиотеке
	// (GET /api/v1/libraries/{libraryUid}/books)
	ListBooks(ctx echo.Context, libraryUid openapi_types.UUID, params ListBooksParams) error
	// Взять книгу в библиотеке
	// (POST /api/v1/libraries/{libraryUid}/books/{bookUid})
	TakeBook(ctx echo.Context, libraryUid openapi_types.UUID, bookUid openapi_types.UUID) error
	// Вернуть книгу в библиотеку
	// (POST /api/v1/libraries/{libraryUid}/books/{bookUid}/return)
	ReturnBook(ctx echo.Context, libraryUid openapi_types.UUID, bookUid openapi_types.UUID) error
	// Проверка живости сервиса
	// (GET /manage/health)
	Health(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetBook converts echo context to params.
func (w *ServerInterfaceWrapper) GetBook(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bookUid" -------------
	var bookUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "bookUid", ctx.Param("bookUid"), &bookUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bookUid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBook(ctx, bookUid)
	return err
}

// ListLibraries converts echo context to params.
func (w *ServerInterfaceWrapper) ListLibraries(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListLibrariesParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// ------------- Required query parameter "city" -------------

	err = runtime.BindQueryParameter("form", true, true, "city", ctx.QueryParams(), &params.City)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter city: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListLibraries(ctx, params)
	return err
}

// GetLibrary converts echo context to params.
func (w *ServerInterfaceWrapper) GetLibrary(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "libraryUid" -------------
	var libraryUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "libraryUid", ctx.Param("libraryUid"), &libraryUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter libraryUid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLibrary(ctx, libraryUid)
	return err
}

// ListBooks converts echo context to params.
func (w *ServerInterfaceWrapper) ListBooks(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "libraryUid" -------------
	var libraryUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "libraryUid", ctx.Param("libraryUid"), &libraryUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter libraryUid: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListBooksParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// ------------- Optional query parameter "showAll" -------------

	err = runtime.BindQueryParameter("form", true, false, "showAll", ctx.QueryParams(), &params.ShowAll)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter showAll: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListBooks(ctx, libraryUid, params)
	return err
}

// TakeBook converts echo context to params.
func (w *ServerInterfaceWrapper) TakeBook(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "libraryUid" -------------
	var libraryUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "libraryUid", ctx.Param("libraryUid"), &libraryUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter libraryUid: %s", err))
	}

	// ------------- Path parameter "bookUid" -------------
	var bookUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "bookUid", ctx.Param("bookUid"), &bookUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bookUid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.TakeBook(ctx, libraryUid, bookUid)
	return err
}

// ReturnBook converts echo context to params.
func (w *ServerInterfaceWrapper) ReturnBook(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "libraryUid" -------------
	var libraryUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "libraryUid", ctx.Param("libraryUid"), &libraryUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter libraryUid: %s", err))
	}

	// ------------- Path parameter "bookUid" -------------
	var bookUid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "bookUid", ctx.Param("bookUid"), &bookUid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bookUid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ReturnBook(ctx, libraryUid, bookUid)
	return err
}

// Health converts echo context to params.
func (w *ServerInterfaceWrapper) Health(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Health(ctx)
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
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/books/:bookUid", wrapper.GetBook)
	router.GET(baseURL+"/api/v1/libraries", wrapper.ListLibraries)
	router.GET(baseURL+"/api/v1/libraries/:libraryUid", wrapper.GetLibrary)
	router.GET(baseURL+"/api/v1/libraries/:libraryUid/books", wrapper.ListBooks)
	router.POST(baseURL+"/api/v1/libraries/:libraryUid/books/:bookUid", wrapper.TakeBook)
	router.POST(baseURL+"/api/v1/libraries/:libraryUid/books/:bookUid/return", wrapper.ReturnBook)
	router.GET(baseURL+"/manage/health", wrapper.Health)

}

type GetBookRequestObject struct {
	BookUid openapi_types.UUID `json:"bookUid"`
}

type GetBookResponseObject interface {
	VisitGetBookResponse(w http.ResponseWriter) error
}

type GetBook200JSONResponse BookInfo

func (response GetBook200JSONResponse) VisitGetBookResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListLibrariesRequestObject struct {
	Params ListLibrariesParams
}

type ListLibrariesResponseObject interface {
	VisitListLibrariesResponse(w http.ResponseWriter) error
}

type ListLibraries200JSONResponse LibraryPaginationResponse

func (response ListLibraries200JSONResponse) VisitListLibrariesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetLibraryRequestObject struct {
	LibraryUid openapi_types.UUID `json:"libraryUid"`
}

type GetLibraryResponseObject interface {
	VisitGetLibraryResponse(w http.ResponseWriter) error
}

type GetLibrary200JSONResponse LibraryResponse

func (response GetLibrary200JSONResponse) VisitGetLibraryResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListBooksRequestObject struct {
	LibraryUid openapi_types.UUID `json:"libraryUid"`
	Params     ListBooksParams
}

type ListBooksResponseObject interface {
	VisitListBooksResponse(w http.ResponseWriter) error
}

type ListBooks200JSONResponse LibraryBookPaginationResponse

func (response ListBooks200JSONResponse) VisitListBooksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type TakeBookRequestObject struct {
	LibraryUid openapi_types.UUID `json:"libraryUid"`
	BookUid    openapi_types.UUID `json:"bookUid"`
}

type TakeBookResponseObject interface {
	VisitTakeBookResponse(w http.ResponseWriter) error
}

type TakeBook204Response struct {
}

func (response TakeBook204Response) VisitTakeBookResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type TakeBook400JSONResponse ValidationErrorResponse

func (response TakeBook400JSONResponse) VisitTakeBookResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReturnBookRequestObject struct {
	LibraryUid openapi_types.UUID `json:"libraryUid"`
	BookUid    openapi_types.UUID `json:"bookUid"`
	Body       *ReturnBookJSONRequestBody
}

type ReturnBookResponseObject interface {
	VisitReturnBookResponse(w http.ResponseWriter) error
}

type ReturnBook200JSONResponse ViolationStatus

func (response ReturnBook200JSONResponse) VisitReturnBookResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReturnBook400JSONResponse ValidationErrorResponse

func (response ReturnBook400JSONResponse) VisitReturnBookResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type HealthRequestObject struct {
}

type HealthResponseObject interface {
	VisitHealthResponse(w http.ResponseWriter) error
}

type Health200Response struct {
}

func (response Health200Response) VisitHealthResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Получить информацию о книге
	// (GET /api/v1/books/{bookUid})
	GetBook(ctx context.Context, request GetBookRequestObject) (GetBookResponseObject, error)
	// Получить список библиотек в городе
	// (GET /api/v1/libraries)
	ListLibraries(ctx context.Context, request ListLibrariesRequestObject) (ListLibrariesResponseObject, error)
	// Получить информацию о библиотеке
	// (GET /api/v1/libraries/{libraryUid})
	GetLibrary(ctx context.Context, request GetLibraryRequestObject) (GetLibraryResponseObject, error)
	// Получить список книг в выбранной библиотеке
	// (GET /api/v1/libraries/{libraryUid}/books)
	ListBooks(ctx context.Context, request ListBooksRequestObject) (ListBooksResponseObject, error)
	// Взять книгу в библиотеке
	// (POST /api/v1/libraries/{libraryUid}/books/{bookUid})
	TakeBook(ctx context.Context, request TakeBookRequestObject) (TakeBookResponseObject, error)
	// Вернуть книгу в библиотеку
	// (POST /api/v1/libraries/{libraryUid}/books/{bookUid}/return)
	ReturnBook(ctx context.Context, request ReturnBookRequestObject) (ReturnBookResponseObject, error)
	// Проверка живости сервиса
	// (GET /manage/health)
	Health(ctx context.Context, request HealthRequestObject) (HealthResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetBook operation middleware
func (sh *strictHandler) GetBook(ctx echo.Context, bookUid openapi_types.UUID) error {
	var request GetBookRequestObject

	request.BookUid = bookUid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetBook(ctx.Request().Context(), request.(GetBookRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBook")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetBookResponseObject); ok {
		return validResponse.VisitGetBookResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ListLibraries operation middleware
func (sh *strictHandler) ListLibraries(ctx echo.Context, params ListLibrariesParams) error {
	var request ListLibrariesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListLibraries(ctx.Request().Context(), request.(ListLibrariesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListLibraries")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ListLibrariesResponseObject); ok {
		return validResponse.VisitListLibrariesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetLibrary operation middleware
func (sh *strictHandler) GetLibrary(ctx echo.Context, libraryUid openapi_types.UUID) error {
	var request GetLibraryRequestObject

	request.LibraryUid = libraryUid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetLibrary(ctx.Request().Context(), request.(GetLibraryRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetLibrary")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetLibraryResponseObject); ok {
		return validResponse.VisitGetLibraryResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ListBooks operation middleware
func (sh *strictHandler) ListBooks(ctx echo.Context, libraryUid openapi_types.UUID, params ListBooksParams) error {
	var request ListBooksRequestObject

	request.LibraryUid = libraryUid
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListBooks(ctx.Request().Context(), request.(ListBooksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListBooks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ListBooksResponseObject); ok {
		return validResponse.VisitListBooksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// TakeBook operation middleware
func (sh *strictHandler) TakeBook(ctx echo.Context, libraryUid openapi_types.UUID, bookUid openapi_types.UUID) error {
	var request TakeBookRequestObject

	request.LibraryUid = libraryUid
	request.BookUid = bookUid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TakeBook(ctx.Request().Context(), request.(TakeBookRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TakeBook")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(TakeBookResponseObject); ok {
		return validResponse.VisitTakeBookResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ReturnBook operation middleware
func (sh *strictHandler) ReturnBook(ctx echo.Context, libraryUid openapi_types.UUID, bookUid openapi_types.UUID) error {
	var request ReturnBookRequestObject

	request.LibraryUid = libraryUid
	request.BookUid = bookUid

	var body ReturnBookJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReturnBook(ctx.Request().Context(), request.(ReturnBookRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReturnBook")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ReturnBookResponseObject); ok {
		return validResponse.VisitReturnBookResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Health operation middleware
func (sh *strictHandler) Health(ctx echo.Context) error {
	var request HealthRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Health(ctx.Request().Context(), request.(HealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Health")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(HealthResponseObject); ok {
		return validResponse.VisitHealthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
