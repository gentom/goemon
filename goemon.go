package goemon

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	Goemon struct {
		router *Router
	}

	Handle func(http.ResponseWriter, *http.Request, url.Values)
)

// HTTP methods
const (
	DELETE = "DELETE"
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
)

func New() (g *Goemon) {
	g = &Goemon{}
	g.router = NewRouter()
	return g
}

func (g *Goemon) GET(path string, handler Handle) {
	g.router.tree.Add(GET, path, handler)
}

func (g *Goemon) POST(path string, handler Handle) {
	g.router.tree.Add(POST, path, handler)
}

func (g *Goemon) PUT(path string, handler Handle) {
	g.router.tree.Add(PUT, path, handler)
}

func (g *Goemon) DELETE(path string, handler Handle) {
	g.router.tree.Add(DELETE, path, handler)
}

func (g *Goemon) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := g.router.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		handler(w, req, params)
	}
}

func (g *Goemon) Start(port int) error {
	portString := fmt.Sprintf(":%d", port)
	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         portString,
		Handler:      g,
	}
	return server.ListenAndServe()
}
