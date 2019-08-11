package handler_test

import (
	. "github.com/egoholic/router/test/helper"

	. "github.com/egoholic/router/handler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var description = "handler desription"
var _ = Describe("handler", func() {
	var handler *Handler
	BeforeEach(func() {
		handler = New(ExampleHandlerFunc, description)
	})
	Describe("New()", func() {
		It("returns handler", func() {
			Expect(handler).NotTo(BeNil())
			Expect(handler).To(BeAssignableToTypeOf(&Handler{}))
		})
	})
	Describe("handler", func() {
		Describe(".HandlerFn()", func() {
			It("returns handler function", func() {
				var hfn HandlerFn
				Expect(handler.HandlerFn()).To(BeAssignableToTypeOf(hfn))
			})
		})
		Describe(".Description()", func() {
			It("returns handler's description", func() {
				Expect(handler.Description()).To(Equal(description))
			})
		})
	})
})
