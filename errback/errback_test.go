package errback_test

import (
	"io/ioutil"
	"net/http"
	. "net/http/httptest"
	"strings"

	. "github.com/egoholic/router/errback"
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/params"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeHandler(msg string) handler.HandlerFn {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		w.Write([]byte(msg))
	}
}

var (
	prms *params.Params
	resp *ResponseRecorder
	req  *http.Request
)

var _ = Describe("errback", func() {
	BeforeEach(func() {
		prms = params.New("/", http.MethodGet, map[string]interface{}{})
		resp = NewRecorder()
		req = NewRequest(http.MethodGet, "/", strings.NewReader("body"))
	})
	Describe("New()", func() {
		Context("when configured properly", func() {
			It("returns ErrBack", func() {
				eb, err := New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(eb).To(BeAssignableToTypeOf(&ErrBack{}))
				Expect(err).To(BeNil())
			})
		})
		Context("when NOT configured properly", func() {
			It("fails", func() {
				_, err := New(
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("NotFound handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("BadRequest handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("ServerError handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
				)
				Expect(err.Error()).To(Equal("Unauthorized handler is not defined"))
				_, err = New(
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("NotFound handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithNotFound(makeHandler(NOT_FOUND)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("BadRequest handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				Expect(err.Error()).To(Equal("ServerError handler is not defined"))
				_, err = New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithServerError(makeHandler(SERVER_ERROR)),
				)
				Expect(err.Error()).To(Equal("Unauthorized handler is not defined"))
			})
		})
	})
	Describe("ErrBack", func() {
		Describe(".HandleNotFound()", func() {
			It("sets correct status code and renders error message", func() {
				eb, err := New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				eb.HandleNotFound(resp, req, prms)
				response := resp.Result()
				Expect(response.StatusCode).To(Equal(NOT_FOUND_CODE))
				body, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte(NOT_FOUND)))
			})
		})
		Describe(".HandleBadRequest()", func() {
			It("sets correct status code and renders error message", func() {
				eb, err := New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				eb.HandleBadRequest(resp, req, prms)
				response := resp.Result()
				Expect(response.StatusCode).To(Equal(BAD_REQUEST_CODE))
				body, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte(BAD_REQUEST)))
			})
		})
		Describe(".HandleServerError()", func() {
			It("sets correct status code and renders error message", func() {
				eb, err := New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				eb.HandleServerError(resp, req, prms)
				response := resp.Result()
				Expect(response.StatusCode).To(Equal(SERVER_ERROR_CODE))
				body, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte(SERVER_ERROR)))
			})
		})
		Describe(".HandlUnauthorized()", func() {
			It("sets correct status code and renders error message", func() {
				eb, err := New(
					WithNotFound(makeHandler(NOT_FOUND)),
					WithBadRequest(makeHandler(BAD_REQUEST)),
					WithServerError(makeHandler(SERVER_ERROR)),
					WithUnauthorized(makeHandler(UNAUTHORIZED)),
				)
				eb.HandleBadRequest(resp, req, prms)
				response := resp.Result()
				Expect(response.StatusCode).To(Equal(BAD_REQUEST_CODE))
				body, err := ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte(BAD_REQUEST)))
			})
		})
	})
})
