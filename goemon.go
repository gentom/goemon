package goemon

import (
	"net/http"
	"net/url"
	"strings"
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

func (n *node) Add(method, path string, handler Handle) {
	if path == "" {
		panic("Path cannot be empty")
	}
	if path[0] != '/' {
		path = "/" + path
	}
	components := strings.Split(path, "/")[1:]
	for count := len(components); count > 0; count-- {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 {
			aNode.methods[method] = handler
			return
		}
		newNode := node{
			component:    component,
			isNamedParam: false,
			methods:      make(map[string]Handle),
		}

		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}
		if count == 1 {
			newNode.methods[method] = handler
		}
		aNode.children = append(aNode.children, &newNode)
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
