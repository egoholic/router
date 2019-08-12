package params_test

import (
	. "github.com/egoholic/router/params"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Params", func() {
	var (
		params  *Params
		path1   = "/articles"
		method1 = "GET"
	)

	Context("creation", func() {
		Describe("New()", func() {
			It("returns params", func() {
				p := New(path1, method1, map[string]interface{}{})
				Expect(p).To(BeAssignableToTypeOf(&Params{}))
				Expect(p.Path()).To(Equal(path1))
				Expect(p.Verb()).To(Equal(method1))
			})
		})
	})
	Context("accessors", func() {
		BeforeEach(func() {
			prmsMap := map[string]interface{}{}
			prmsMap[":article_id"] = 11
			params = New(path1, method1, prmsMap)
		})
		Describe(".Path()", func() {
			It("returns path", func() {
				Expect(params.Path()).To(Equal(path1))
			})
		})
		Describe(".PathChunks()", func() {
			It("returns path", func() {
				Expect(params.PathChunks()).To(Equal([]string{"", "articles"}))
			})
		})
		Describe(".Verb()", func() {
			It("returns path", func() {
				Expect(params.Verb()).To(Equal(method1))
			})
		})
		Describe(".Get()", func() {
			Context("when param exists", func() {
				It("returns param's value and ok", func() {
					val, ok := params.Get(":article_id")
					Expect(ok).To(BeTrue())
					Expect(val).NotTo(BeNil())
					Expect(val.(int)).To(Equal(11))
				})
			})
			Context("when there is no param with given name", func() {
				It("returns nil and not ok", func() {
					val, ok := params.Get(":rubric_id")
					Expect(ok).To(BeFalse())
					Expect(val).To(BeNil())
				})
			})
		})
		Describe(".Set()", func() {
			Context("when param exists", func() {
				It("updates value", func() {
					val, ok := params.Get(":article_id")
					Expect(ok).To(BeTrue())
					Expect(val).NotTo(BeNil())
					Expect(val.(int)).To(Equal(11))
					params.Set(":article_id", 13)
					val, ok = params.Get(":article_id")
					Expect(ok).To(BeTrue())
					Expect(val).NotTo(BeNil())
					Expect(val.(int)).To(Equal(13))
				})
			})
			Context("when there is no param with given name", func() {
				It("sets value", func() {
					val, ok := params.Get(":rubric_id")
					Expect(ok).To(BeFalse())
					Expect(val).To(BeNil())
					params.Set(":rubric_id", 3)
					val, ok = params.Get(":rubric_id")
					Expect(ok).To(BeTrue())
					Expect(val).NotTo(BeNil())
					Expect(val.(int)).To(Equal(3))
				})
			})
		})
	})
	Describe(".NewIterator()", func() {
		It("returns iterator", func() {
			params = New(path1, method1, map[string]interface{}{})
			Expect(params.NewIterator()).To(BeAssignableToTypeOf(&PathChunksIterator{}))
		})
	})
	Describe("PathChunksIterator", func() {
		var iterator *PathChunksIterator
		BeforeEach(func() {
			params = New(path1, method1, map[string]interface{}{})
			iterator = params.NewIterator()
		})
		Describe(".HasNext()", func() {
			Context("when has next", func() {
				It("returns true", func() {
					Expect(iterator.Current()).To(Equal(""))
					Expect(iterator.HasNext()).To(BeTrue())
				})
			})
			Context("when has no next", func() {
				It("returns false", func() {
					Expect(iterator.Current()).To(Equal(""))
					iterator.Next()
					Expect(iterator.Current()).To(Equal("articles"))
					Expect(iterator.HasNext()).To(BeFalse())
				})
			})
		})
		Describe(".Next()", func() {
			Context("when has next", func() {
				It("returns path chunk", func() {
					chunk, _ := iterator.Next()
					Expect(chunk).To(Equal("articles"))
				})
			})
		})
		Describe(".Current()", func() {
			It("returns current path chunk", func() {
				Expect(iterator.Current()).To(Equal(""))
			})
		})
	})
})
