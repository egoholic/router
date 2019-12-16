package node

import (
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/path/chunk"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type (
	Node struct {
		chunk    *chunk.Chunk
		children []*Node
		handlers map[string]*handler.Handler
	}
	Option func(*Node)
)

func New(pathChunk *chunk.Chunk, opts ...Option) *Node {
	node := &Node{
		chunk: pathChunk,
	}
	for _, opt := range opts {
		opt(node)
	}
	return node
}
func WithNewChild(pathChunk *chunk.Chunk, opts ...Option) Option {
	child := New(pathChunk, opts...)
	return func(n *Node) {
		n.children = append(n.children, child)
	}
}
func WithChild(child *Node) Option {
	return func(n *Node) {
		n.children = append(n.children, child)
	}
}
func WithGetHandler(h *handler.Handler) Option {
	return func(n *Node) {
		n.handlers[GET] = h
	}
}
func WithPostHandler(h *handler.Handler) Option {
	return func(n *Node) {
		n.handlers[POST] = h
	}
}
func WithPutHandler(h *handler.Handler) Option {
	return func(n *Node) {
		n.handlers[PUT] = h
	}
}
func WithPatchHandler(h *handler.Handler) Option {
	return func(n *Node) {
		n.handlers[PATCH] = h
	}
}
func WithDeleteHandler(h *handler.Handler) Option {
	return func(n *Node) {
		n.handlers[DELETE] = h
	}
}
