package goemon

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func test(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "TEST!\n")
}

func TestGETReq(t *testing.T) {

	api := New()
	api.GET("/test", test)
	go api.Start(3000)

	resp, err := http.Get("http://localhost:3000/test")
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "TEST!\n" {
		t.Error("Not equal.")
	}
}
