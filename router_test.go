package router

import (
	"fmt"
	"net/http"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func TestRouter(t *testing.T) {
	h := http.HandlerFunc(hello)
	root := CreateRouteNode()
	root.Insert("/user/<int:id>", []string{"GET"}, h)
	root.Insert("/user/<int:id>/<string:action>/wuliwal", []string{"POST"}, h)
	root.Insert("/user/caca<int:id>/tutu<string:action>asd$", []string{"POST"}, h)
	matched, _, params := root.Match("/user/caca123/tutudeleteasd$", "GET")
	if matched {
		for _, v := range params {
			t.Log(v.Name, v.Value)
		}
		t.Error("must be not matched")
	}
	matched, _, params = root.Match("/user/caca123/tutudeleteasd$", "POST")
	if matched {
		for _, v := range params {
			t.Log(v.Name, v.Value)
		}
	} else {
		t.Error("must be matched")
	}

}
