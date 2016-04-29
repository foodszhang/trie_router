package trieRouter

import (
	"net/http"
	"strings"
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
	reg         string
	handler     http.HandlerFunc
	middlewares []Adapter
}

type Adapter func(http.Handler) http.Handler

// CreateRouteNode init a route node
func CreateRouteNode() *RouteNode {
	node := new(RouteNode)
	node.childNodes = make(map[rune]*RouteNode)
	return node
}

func getPrefixReg(url string) (string, string) {
	url = strings.Replace(url, " ", "", -1)
	lastStop := 0
	for i, c := range url {
		if c == '/' {
			lastStop = i
		}
		if c == '<' || c == '>' {
			break
		}
	}
	return url[0:lastStop], url[lastStop:]

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

// Search 函数接收一个给定字符串, 返回一个Route
