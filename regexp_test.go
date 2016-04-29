package trieRouter

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	r := regexp.MustCompile("<.*$")
	all := r.FindString("/user/<int:id>")
	t.Log(all)
	all = r.FindString("/user/<int:id>/<string:action>/wuliwal")
	t.Log(all)
	all = r.FindString("/user/caca<int:id>/tutu<string:action>")
	t.Log(all)
	all = r.FindString("/user/caca<int:id>/tutu<string:action>asd")
	t.Log(all)
	sp := r.Split("/user/caca<int:id>/tutu<string:action>asd$", -1)
	t.Log(sp)
}
