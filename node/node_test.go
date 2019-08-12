package node_test

import (
	"strconv"

	rtr "github.com/egoholic/router"
	. "github.com/egoholic/router/node"
	"github.com/egoholic/router/params"
	. "github.com/egoholic/router/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type articleIDForm struct{}

func (f *articleIDForm) CheckAndPopulate(pattern string, chunk string, prms *params.Params) bool {
	num, err := strconv.ParseInt(chunk, 10, 64)
	if err != nil {
		return false
	}
	prms.Set(pattern, num)
	return true
}

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
		Describe(".Child()", func() {
			Context("when node with given path chunk exists", func() {
				It("returns existing Node", func() {
					node2 := root.Child("articles", &DumbForm{})
					node3 := root.Child("articles", &DumbForm{})
					Expect(node3).To(BeIdenticalTo(node2))
				})
			})
			Context("when node with given path chunk not exists", func() {
				It("creates new Node and returns it", func() {
					Expect(root.Child("articles", &DumbForm{})).To(BeAssignableToTypeOf(&Node{}))
				})
			})
		})
		Describe(".GET()", func() {
			It("defines GET request handler", func() {
				prms := params.New("/articles", "GET", map[string]interface{}{})
				Expect(router.Handler(prms)).To(BeNil())
				articles := root.Child("articles", &DumbForm{})
				articles.GET(ExampleHandlerFunc, "description1")
				Expect(router.Handler(prms)).NotTo(BeNil())

				prms = params.New("/articles/11", "GET", map[string]interface{}{})
				article := articles.Child(":article_id", &articleIDForm{})
				article.GET(ExampleHandlerFunc, "description2")
				Expect(router.Handler(prms)).NotTo(BeNil())
				val, ok := prms.Get(":article_id")
				Expect(ok).To(BeTrue())
				Expect(val).NotTo(BeNil())
				Expect(val.(int64)).To(Equal(int64(11)))
			})
		})
		Describe(".POST()", func() {
			It("defines POST request handler", func() {
				params := params.New("/articles", "POST", map[string]interface{}{})
				Expect(router.Handler(params)).To(BeNil())
				root.Child("articles", &DumbForm{}).POST(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".PUT()", func() {
			It("defines PUT request handler", func() {
				params := params.New("/", "PUT", map[string]interface{}{})
				Expect(router.Handler(params)).To(BeNil())
				root.PUT(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".PATCH()", func() {
			It("defines PATCH request handler", func() {
				params := params.New("/", "PATCH", map[string]interface{}{})
				Expect(router.Handler(params)).To(BeNil())
				root.PATCH(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
		Describe(".DELETE()", func() {
			It("defines DELETE request handler", func() {
				params := params.New("/articles", "DELETE", map[string]interface{}{})
				Expect(router.Handler(params)).To(BeNil())
				root.Child("articles", &DumbForm{}).DELETE(ExampleHandlerFunc, "description1")
				Expect(router.Handler(params)).NotTo(BeNil())
			})
		})
	})
})
