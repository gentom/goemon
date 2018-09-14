# Goemon  
Goemon is a micro framework which is zero-dependencies for building APIs.    
Goemon is still in development.

## Usage
```
package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gentom/goemon"
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
```
If you run the code above and curl the endpoint:
```
$ curl http://localhost:5000/
root

$ curl http://localhost:5000/hello
hello world!

$ curl http://localhost:5000/langs/nim
[nim] is awesome!

$ curl http://localhost:5000/langs/golang/framework
This is [golang]'s framework
```