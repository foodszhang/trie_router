package trieRouter

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
	root.Insert("/user/<int:id>", h)
	root.Insert("/user/<int:id>/<string:action>/wuliwal", h)
	root.Insert("/user/caca<int:id>/tutu<string:action>asd$", h)
	matched, _, params := root.Match("/user/caca123/tutudeleteasd$")
	if matched {
		for _, v := range params {
			t.Log(v.name, v.value)
		}
	}

}
