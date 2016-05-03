package trieRouter

import (
	"regexp"
)

type Param struct {
	name  string
	reg   *regexp.Regexp
	value string
}

var rReg *regexp.Regexp = regexp.MustCompile(`<.*$`)
var pReg *regexp.Regexp = regexp.MustCompile(`<([^<>:]*):([^<>:]*)>`)
var typeReg map[string]*regexp.Regexp

func getPrefixReg(url string) (string, string) {
	reg := rReg.FindString(url)
	prefix := rReg.Split(url, -1)[0]
	return prefix, reg
}

//获取<>中的类型信息,返回应该有的参数列表
func getRegType(reg string) []Param {
	all := pReg.FindAllStringSubmatch(reg, -1)
	params := make([]Param, len(all))
	for i := range all {
		params[i].name = all[i][1]
		params[i].reg = typeReg[all[i][1]]
	}
	return params
}

//将reg中的<>替换为相应的正则表达并
func replaceReg(reg string) string {
	return pReg.ReplaceAllStringFunc(reg, rg)

}
func rg(reg string) string {
	param := getRegType(reg)[0]
	return param.reg.String()

}
func init() {
	typeReg["int"] = regexp.MustCompile(`\d+`)
	typeReg["string"] = regexp.MustCompile(`[^/]+`)
}
