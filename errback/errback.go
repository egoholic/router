package errback

import (
	"errors"
	"net/http"

	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/params"
)

const (
	BAD_REQUEST_CODE  = 400
	UNAUTHORIZED_CODE = 401
	NOT_FOUND_CODE    = 404
	SERVER_ERROR_CODE = 500

	BAD_REQUEST  = "Bad Request (400)"
	UNAUTHORIZED = "Unauthorized (401)"
	NOT_FOUND    = "Not Found (404)"
	SERVER_ERROR = "Internal Server Error (500)"
)

type ErrBack struct {
	notFound     handler.HandlerFn
	badRequest   handler.HandlerFn
	serverError  handler.HandlerFn
	unauthorized handler.HandlerFn
}

type Conf func(eb *ErrBack)

func WithNotFound(h handler.HandlerFn) Conf {
	return func(eb *ErrBack) {
		(*eb).notFound = h
	}
}
func WithBadRequest(h handler.HandlerFn) Conf {
	return func(eb *ErrBack) {
		(*eb).badRequest = h
	}
}
func WithServerError(h handler.HandlerFn) Conf {
	return func(eb *ErrBack) {
		(*eb).serverError = h
	}
}
func WithUnauthorized(h handler.HandlerFn) Conf {
	return func(eb *ErrBack) {
		(*eb).unauthorized = h
	}
}
func New(configs ...Conf) (*ErrBack, error) {
	var (
		err error
		eb  = &ErrBack{}
	)
	for _, cfg := range configs {
		cfg(eb)
	}
	if eb.notFound == nil {
		err = errors.New("NotFound handler is not defined")
		return eb, err
	}
	if eb.badRequest == nil {
		err = errors.New("BadRequest handler is not defined")
		return eb, err
	}
	if eb.serverError == nil {
		err = errors.New("ServerError handler is not defined")
		return eb, err
	}
	if eb.unauthorized == nil {
		err = errors.New("Unauthorized handler is not defined")
		return eb, err
	}
	return eb, err
}

func (eb *ErrBack) HandleNotFound(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(NOT_FOUND_CODE)
	eb.notFound(w, r, p)
}
func (eb *ErrBack) HandleBadRequest(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(BAD_REQUEST_CODE)
	eb.badRequest(w, r, p)
}
func (eb *ErrBack) HandleServerError(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(SERVER_ERROR_CODE)
	eb.serverError(w, r, p)
}
func (eb *ErrBack) HandleUnauthorized(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(UNAUTHORIZED_CODE)
	eb.unauthorized(w, r, p)
}
