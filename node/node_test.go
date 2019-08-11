package node_test

import (
	rtr "github.com/egoholic/router"
	. "github.com/egoholic/router/node"
	"github.com/egoholic/router/params"
	. "github.com/egoholic/router/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("node", func() {
	var (
		router *rtr.Router
		root   *Node
	)
	BeforeEach(func() {
		router = rtr.New()
		root = router.Root()
	})
	Describe("New()", func() {
		It("returns node", func() {
			Expect(root).NotTo(BeNil())
			Expect(root).To(BeAssignableToTypeOf(&Node{}))
		})
	})
	Describe("Node", func() {
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
				root.Sub("articles").GET(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".POST()", func() {
			It("defines POST request handler", func() {
				params := params.New("/articles", "POST", map[string][]string{})
				Expect(router.Handler(params)).To(BeNil())
				root.Sub("articles").POST(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".PUT()", func() {
			It("defines PUT request handler", func() {
				params := params.New("/", "PUT", map[string][]string{})
				Expect(router.Handler(params)).To(BeNil())
				root.PUT(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".PATCH()", func() {
			It("defines PATCH request handler", func() {
				params := params.New("/", "PATCH", map[string][]string{})
				Expect(router.Handler(params)).To(BeNil())
				root.PATCH(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".DELETE()", func() {
			It("defines DELETE request handler", func() {
				params := params.New("/articles", "DELETE", map[string][]string{})
				Expect(router.Handler(params)).To(BeNil())
				root.Sub("articles").DELETE(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
	})
})
