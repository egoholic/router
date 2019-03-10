package router_test

import (
	. "github.com/egoholic/router"
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
				Expect(r.Root()).To(BeAssignableToTypeOf(&Node{}))
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
					prms := params.New("/", GET, _prms)
					r.Root().GET(TestHandlerFunc, description)
					handler := r.Handler(prms)
					Expect(handler).To(BeAssignableToTypeOf(&Handler{}))

					_prms2 := map[string][]string{}
					_prms2["header"] = []string{"TestHeader2"}
					description2 := "description2"
					prms2 := params.New("/articles", GET, _prms2)

					r.Root().Sub("articles").GET(TestHandlerFunc, description2)
					handler = r.Handler(prms2)
					Expect(handler).To(BeAssignableToTypeOf(&Handler{}))
				})
			})

			Context("when route does not exist", func() {
				It("returns nil", func() {
					r := New()
					params := params.New("/", GET, map[string][]string{})
					Expect(r.Handler(params)).To(BeNil())
				})
			})
		})
	})

	Describe("Node", func() {
		var (
			router *Router
			root   *Node
		)

		BeforeEach(func() {
			router = New()
			root = router.Root()
		})

		Context("definition", func() {

			Describe(".Sub()", func() {
				Context("when node with given path chunk exists", func() {
					It("returns existing Node", func() {
						node2 := root.Sub("articles")
						node3 := root.Sub("articles")
						Expect(node3).To(BeIdenticalTo(node2))
					})
				})

				Context("when node with given path chunk not exists", func() {
					It("creates new Node and returns it", func() {
						Expect(root.Sub("articles")).To(BeAssignableToTypeOf(&Node{}))
					})
				})
			})

			Describe(".GET()", func() {
				It("defines GET request handler", func() {
					params := params.New("/articles", "GET", map[string][]string{})
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").GET(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".POST()", func() {
				It("defines POST request handler", func() {
					params := params.New("/articles", "POST", map[string][]string{})
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").POST(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PUT()", func() {
				It("defines PUT request handler", func() {
					params := params.New("/", "PUT", map[string][]string{})
					Expect(router.Handler(params)).To(BeNil())
					root.PUT(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PATCH()", func() {
				It("defines PATCH request handler", func() {
					params := params.New("/", "PATCH", map[string][]string{})
					Expect(router.Handler(params)).To(BeNil())
					root.PATCH(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".DELETE()", func() {
				It("defines DELETE request handler", func() {
					params := params.New("/articles", "DELETE", map[string][]string{})
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").DELETE(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})
		})
	})
})
