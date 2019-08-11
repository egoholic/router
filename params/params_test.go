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
				p := New(path1, method1, map[string][]string{})
				Expect(p).To(BeAssignableToTypeOf(&Params{}))
				Expect(p.Path()).To(Equal(path1))
				Expect(p.Verb()).To(Equal(method1))
			})
		})
	})
	Context("accessors", func() {
		BeforeEach(func() {
			params = New(path1, method1, map[string][]string{})
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
	})
	Describe(".NewIterator()", func() {
		It("returns iterator", func() {
			params = New(path1, method1, map[string][]string{})
			Expect(params.NewIterator()).To(BeAssignableToTypeOf(&PathChunksIterator{}))
		})
	})
	Describe("PathChunksIterator", func() {
		var iterator *PathChunksIterator
		BeforeEach(func() {
			params = New(path1, method1, map[string][]string{})
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
