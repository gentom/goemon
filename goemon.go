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
