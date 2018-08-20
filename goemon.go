package goemon

import (
	"net/http"
	"net/url"
)

type (
	GOEMON struct {
		router *Router
	}

	Router struct {
		tree *node
	}

	node struct {
		children     []*node
		component    string
		isNamedParam bool
		methods      map[string]Handle
	}

	Handle func(http.ResponseWriter, *http.Request, url.Values)
)

func New() (g *GOEMON) {
	g = &GOEMON{}
	g.router = NewRouter()
	return g
}

func NewRouter() *Router {
	node := node{
		component:    "/",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}
	return &Router{tree: &node}
}
