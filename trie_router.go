package trieRouter

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
	Handler     http.HandlerFunc
	Middlewares []Adapter
}

type Adapter func(http.Handler) http.Handler

// Adapt wrap all adaters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for i := range adapters {
		h = adapters[len(adapters)-i-1](h)
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

// Insert insert a url with handle
// Insert 应当接收一个url字串和一个绑定方法, 一些中间件的操作
func (root *RouteNode) Insert(s string, handler http.HandlerFunc, adapters ...Adapter) {
	node := root
	prefix, reg := getPrefixReg(s)
	for _, r := range prefix {
		if node.childNodes[r] == nil {
			node.childNodes[r] = CreateRouteNode()
		}
		node = node.childNodes[r]
	}
	route := Route{reg, handler, adapters}
	node.Exist = true
	node.Routes = append(node.Routes, route)
}
func (root *RouteNode) Match(url string) (bool, http.Handler, []Param) {
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
			if v.Reg == "" {
				return true, Adapt(v.Handler, v.Middlewares...), nil
			}

		}
		return false, nil, nil
	} else if node == nil && lastnode.Exist {
		urlRg := url[len+1:]
		for _, v := range lastnode.Routes {
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
