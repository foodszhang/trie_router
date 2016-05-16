package router

import (
	"net/http"
)

//RouteNode store the route data and be a tree node
//一个前缀节点上可能包括多个需要匹配的路由
type RouteNode struct {
	childNodes map[rune]*RouteNode
	Char       rune
	Exist      bool
	Routes     []Route
}

//每个路由应当只绑定一个方法,但是可以绑定多个中间件
type Route struct {
	Reg         string
	Handler     http.Handler
	Methods     map[string]bool
	Middlewares []Adapter
}

type Adapter interface {
	Adapt(http.Handler) http.Handler
}

// Adapt wrap all adaters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for i := range adapters {
		h = adapters[len(adapters)-i-1].Adapt(h)
	}
	return h

}

// CreateRouteNode init a route node
func CreateRouteNode() *RouteNode {
	node := new(RouteNode)
	node.childNodes = make(map[rune]*RouteNode)
	node.Routes = make([]Route, 0)
	return node
}
func setMethods(methods []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range methods {
		m[v] = true
	}
	return m
}

// Insert insert a url with handle
// Insert 应当接收一个url字串和一个绑定方法,还有一个操作数组, 一些中间件的操作
func (root *RouteNode) Insert(s string, methods []string, handler http.Handler, adapters ...Adapter) error {
	node := root
	prefix, reg := getPrefixReg(s)
	for _, r := range prefix {
		if node.childNodes[r] == nil {
			node.childNodes[r] = CreateRouteNode()
		}
		node = node.childNodes[r]
	}
	route := Route{reg, handler, setMethods(methods), adapters}
	node.Exist = true
	node.Routes = append(node.Routes, route)
	return nil
}
func (root *RouteNode) Match(url, method string) (bool, http.Handler, []Param) {
	node := root
	var lastnode *RouteNode = nil
	len := 0
	for i, r := range url {
		lastnode = node
		node = node.childNodes[r]
		if node == nil {
			break
		}
		len = i
	}
	// 完全匹配的情况,这时候搜索route中reg为空的即可
	if node != nil && node.Exist {
		for _, v := range node.Routes {
			if _, ok := v.Methods[method]; v.Reg == "" && ok {
				return true, Adapt(v.Handler, v.Middlewares...), nil
			}

		}
		return false, nil, nil
	} else if node == nil && lastnode.Exist {
		urlRg := url[len+1:]
		for _, v := range lastnode.Routes {
			if _, ok := v.Methods[method]; !ok {
				continue
			}
			matched, params := Match(v.Reg, urlRg)
			if matched {
				return true, Adapt(v.Handler, v.Middlewares...), params
			}
		}
		return false, nil, nil
	}
	return false, nil, nil
}

// Search 函数接收一个给定字符串, 返回一个Route
