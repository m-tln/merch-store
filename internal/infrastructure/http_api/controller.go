package httpapi

import (
	"encoding/json"
	"net/http"
	"strings"

	openapi "merch-store/api/generated/go"
	"merch-store/pkg/middleware"

	"github.com/gorilla/mux"
)

// CustomAPIController binds http requests to an api service and writes the service results to the http response
type CustomAPIController struct {
	service          CustomAPIService
	errorHandler     openapi.ErrorHandler
	tokenValidateSvc middleware.AuthMiddlewareConfig
}

// CustomAPIOption for how the controller is set up.
type CustomAPIOption func(*CustomAPIController)

// WithCustomAPIErrorHandler inject ErrorHandler into controller
func WithCustomAPIErrorHandler(h openapi.ErrorHandler) CustomAPIOption {
	return func(c *CustomAPIController) {
		c.errorHandler = h
	}
}

// NewCustomAPIController creates a default api controller
func NewCustomAPIController(s CustomAPIService,
	svc middleware.AuthMiddlewareConfig, opts ...CustomAPIOption) *CustomAPIController {
	controller := &CustomAPIController{
		service:          s,
		errorHandler:     openapi.DefaultErrorHandler,
		tokenValidateSvc: svc,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CustomAPIController
func (c *CustomAPIController) Routes() openapi.Routes {

	authMiddleware := middleware.AuthMiddleware(c.tokenValidateSvc)

	return openapi.Routes{
		"ApiInfoGet": openapi.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/api/info",
			HandlerFunc: authMiddleware(c.APIInfoGet),
		},
		"ApiSendCoinPost": openapi.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/sendCoin",
			HandlerFunc: authMiddleware(c.APISendCoinPost),
		},
		"ApiBuyItemGet": openapi.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/api/buy/{item}",
			HandlerFunc: authMiddleware(c.APIBuyItemGet),
		},
		"ApiAuthPost": openapi.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/auth",
			HandlerFunc: c.APIAuthPost,
		},
	}
}

// ApiInfoGet - Получить информацию о монетах, инвентаре и истории транзакций.
func (c *CustomAPIController) APIInfoGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.APIInfoGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = openapi.EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *CustomAPIController) HandlePost(w http.ResponseWriter, r *http.Request) {
	var bodyParam openapi.SendCoinRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &openapi.ParsingError{Err: err}, nil)
		return
	}
	if err := openapi.AssertSendCoinRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := openapi.AssertSendCoinRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.APISendCoinPost(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = openapi.EncodeJSONResponse(result.Body, &result.Code, w)
}

// ApiSendCoinPost - Отправить монеты другому пользователю.
func (c *CustomAPIController) APISendCoinPost(w http.ResponseWriter, r *http.Request) {
	c.HandlePost(w, r)
}

// ApiBuyItemGet - Купить предмет за монеты.
func (c *CustomAPIController) APIBuyItemGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemParam := params["item"]
	if itemParam == "" {
		c.errorHandler(w, r, &openapi.RequiredError{Field: "item"}, nil)
		return
	}
	result, err := c.service.APIBuyItemGet(r.Context(), itemParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = openapi.EncodeJSONResponse(result.Body, &result.Code, w)
}

// ApiAuthPost - Аутентификация и получение JWT-токена.
func (c *CustomAPIController) APIAuthPost(w http.ResponseWriter, r *http.Request) {
	c.HandlePost(w, r)
}
