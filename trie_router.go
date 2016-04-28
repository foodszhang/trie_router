package trieRouter

import (
	"net/http"
	"strings"
)

//RouteNode store the route data and be a tree node
type RouteNode struct {
	childNodes map[rune]*RouteNode
	Char       rune
	Exist      bool
	regxp      []string
	handlers   []http.HandlerFunc
}

func CreateRouteNode() *RouteNode {
	node := new(RouteNode)
	node.childNodes = make(map[rune]*RouteNode)
	return node
}

func getPrefixReg(url string) (string, string) {
	url = strings.Replace(url, " ", "", -1)
	urlbyte := []byte(url)
	lastStop := 0
	for i, c := range urlbyte {
		if c == '/' {
			lastStop = i
		}
		if c == '<' {
			break
		}
	}
	return string(urlbyte[0:lastStop]), string(urlbyte[lastStop:])

}

// Insert insert a url with handle
func (root *RouteNode) Insert(s string) {
	node := root
	for _, r := range s {
		if node.childNodes[r] == nil {
			node.childNodes[r] = CreateRouteNode()
		}
		node = node.childNodes[r]
	}
	node.Exist = true
}
