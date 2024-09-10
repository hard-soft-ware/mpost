package generator

import (
	"regexp"
	"strings"
	"unicode"
)

//###########################################################//

func (n *GeneratorNameObj) ToTitleCase(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	r := regexp.MustCompile(re)
	words := strings.Fields(s)

	for i, word := range words {
		if len(word) > 0 {
			word = r.ReplaceAllString(word, "")
			if len(word) > 0 {
				words[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
			}

		}
	}

	return strings.Join(words, "")
}

////

func (n *GeneratorNameObj) Get() string {
	return n.ToTitleCase(n.gen.name)
}

func (n *GeneratorNameObj) GetParam(param string) string {
	return n.Get() + n.ToTitleCase(param)
}

func (n *GeneratorNameObj) GetObj() string {
	return n.GetParam("Obj")
}

func (n *GeneratorNameObj) GetMap() string {
	return n.GetParam("Map")
}

func (n *GeneratorNameObj) GetType() string {
	return n.GetParam("Type")
}

func (n *GeneratorNameObj) GetText() string {
	return n.GetParam("Text")
}

//

func (n *GeneratorNameObj) GetCode(code string) string {
	return n.Get() + n.ToTitleCase(code)
}

func (n *GeneratorNameObj) GetParamCode(param string, code string) string {
	return n.GetParam(param) + n.ToTitleCase(code)
}

func (n *GeneratorNameObj) GetObjCode(code string) string {
	return n.GetObj() + n.ToTitleCase(code)
}

func (n *GeneratorNameObj) GetMapCode(code string) string {
	return n.GetMap() + n.ToTitleCase(code)
}

func (n *GeneratorNameObj) GetTypeCode(code string) string {
	return n.GetType() + n.ToTitleCase(code)
}

func (n *GeneratorNameObj) GetTextCode(code string) string {
	return n.GetText() + n.ToTitleCase(code)
}

//

func (n *GeneratorNameObj) Self() *GeneratorObj {
	n.gen.Print(n.Get())
	return n.gen
}

func (n *GeneratorNameObj) Param(param string) *GeneratorObj {
	n.gen.Print(n.GetParam(param))
	return n.gen
}

func (n *GeneratorNameObj) Obj() *GeneratorObj {
	n.gen.Print(n.GetObj())
	return n.gen
}

func (n *GeneratorNameObj) Map() *GeneratorObj {
	n.gen.Print(n.GetMap())
	return n.gen
}

func (n *GeneratorNameObj) Type() *GeneratorObj {
	n.gen.Print(n.GetType())
	return n.gen
}

func (n *GeneratorNameObj) Text() *GeneratorObj {
	n.gen.Print(n.GetText())
	return n.gen
}

//

func (n *GeneratorNameObj) CodeToTitleCase(code string) *GeneratorObj {
	n.gen.Print(n.ToTitleCase(code))
	return n.gen
}

func (n *GeneratorNameObj) SelfCode(code string) *GeneratorObj {
	n.gen.Print(n.GetCode(code))
	return n.gen
}

func (n *GeneratorNameObj) SelfParam(param string, code string) *GeneratorObj {
	n.gen.Print(n.GetParamCode(param, code))
	return n.gen
}

func (n *GeneratorNameObj) SelfParamCode(param string, code string) *GeneratorObj {
	n.gen.Print(n.ToTitleCase(param) + n.ToTitleCase(code))
	return n.gen
}

func (n *GeneratorNameObj) ObjCode(code string) *GeneratorObj {
	n.gen.Print(n.GetObjCode(code))
	return n.gen
}

func (n *GeneratorNameObj) MapCode(code string) *GeneratorObj {
	n.gen.Print(n.GetMapCode(code))
	return n.gen
}

func (n *GeneratorNameObj) TypeCode(code string) *GeneratorObj {
	n.gen.Print(n.GetTypeCode(code))
	return n.gen
}

func (n *GeneratorNameObj) TextCode(code string) *GeneratorObj {
	n.gen.Print(n.GetTextCode(code))
	return n.gen
}
