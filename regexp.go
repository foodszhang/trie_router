package trieRouter

import (
	"regexp"
)

var rReg *regexp.Regexp = regexp.MustCompile("<.*$")

func getPrefixReg(url string) (string, string) {
	reg := rReg.FindString(url)
	prefix := rReg.Split(url, -1)[0]
	return reg, prefix
}

//将reg中的<>替换为相应的正则表达并分组:q
