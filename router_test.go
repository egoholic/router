package router_test

import (
	. "github.com/egoholic/router"
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/node"
	"github.com/egoholic/router/params"
	. "github.com/egoholic/router/test/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns router", func() {
				Expect(New()).To(BeAssignableToTypeOf(&Router{}))
			})
		})
	})
	Context("accessors", func() {
		Describe(".Root()", func() {
			It("returns root node", func() {
				r := New()
				Expect(r.Root()).To(BeAssignableToTypeOf(&node.Node{}))
			})
		})
	})
	Context("routing", func() {
		Describe(".Handler()", func() {
			Context("when route exists", func() {
				It("returns handler", func() {
					r := New()
					description := "description"
					_prms := map[string][]string{}
					_prms["header"] = []string{"TestHeader"}
					prms := params.New("/", node.GET, _prms)
					r.Root().GET(ExampleHandlerFunc, description)
					h := r.Handler(prms)
					Expect(h).To(BeAssignableToTypeOf(&handler.Handler{}))

					_prms2 := map[string][]string{}
					_prms2["header"] = []string{"TestHeader2"}
					description2 := "description2"
					prms2 := params.New("/articles", node.GET, _prms2)

					r.Root().Sub("articles").GET(ExampleHandlerFunc, description2)
					h = r.Handler(prms2)
					Expect(h).To(BeAssignableToTypeOf(&handler.Handler{}))
				})
			})
			Context("when route does not exist", func() {
				It("returns nil", func() {
					r := New()
					params := params.New("/", node.GET, map[string][]string{})
					Expect(r.Handler(params)).To(BeNil())
				})
			})
		})
	})
})
