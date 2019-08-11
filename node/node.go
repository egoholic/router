package node

import (
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/params"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Node struct {
	pathChunk    string
	children     map[string]*Node
	verbHandlers map[string]*handler.Handler
}

func New(chunk string) *Node {
	return &Node{chunk, map[string]*Node{}, map[string]*handler.Handler{}}
}
func (n *Node) Sub(pathChunk string) *Node {
	var node *Node
	node = n.children[pathChunk]
	if node != nil {
		return node
	}

	node = New(pathChunk)
	n.children[pathChunk] = node
	return node
}
func (n *Node) Handler(p *params.Params, pathChunks *params.PathChunksIterator) *handler.Handler {
	if pathChunks.HasNext() {
		chunk, _ := pathChunks.Next()
		if child, ok := n.children[chunk]; ok {
			return child.Handler(p, pathChunks)
		}
		return nil
	}
	return n.verbHandlers[p.Verb()]
}
func (n *Node) GET(fn handler.HandlerFn, d string) {
	n.verbHandlers[GET] = handler.New(fn, d)
}
func (n *Node) POST(fn handler.HandlerFn, d string) {
	n.verbHandlers[POST] = handler.New(fn, d)
}
func (n *Node) PUT(fn handler.HandlerFn, d string) {
	n.verbHandlers[PUT] = handler.New(fn, d)
}
func (n *Node) PATCH(fn handler.HandlerFn, d string) {
	n.verbHandlers[PATCH] = handler.New(fn, d)
}
func (n *Node) DELETE(fn handler.HandlerFn, d string) {
	n.verbHandlers[DELETE] = handler.New(fn, d)
}
