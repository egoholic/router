package handler

import (
	"net/http"

	"github.com/egoholic/router/params"
)

type (
	HandlerFn func(w http.ResponseWriter, r *http.Request, p *params.Params)

	Form interface {
		FillAndVerifyParams(string, string, *params.Params)
	}

	Handler struct {
		HandlerFn   HandlerFn
		Form        Form
		Description string
	}
)

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request, p *params.Params) {
	h.Form.FillAndVerifyParams("","",p)
	h.HandlerFn(w, r, p)
}
