package main

import (
	"fmt"
	"net/http"
	"net/url"

	goemon "github.com/gentom/Goemon"
)

func main() {
	api := goemon.New()
	api.GET("/", root)
	api.GET("/hello", hello)
	api.GET("/langs/:lang", programmingLangs)
	api.GET("/langs/:lang/framework", framework)
	api.Start(5000)
}

func root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "root\n")
}

func hello(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "hello world!\n")
}

func programmingLangs(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "%s is awesome!", params["lang"])
}

func framework(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "This is %s's framework", params["lang"])
}
