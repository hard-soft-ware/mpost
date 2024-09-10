package generator

import "bytes"

//###########################################################//

type GeneratorNameObj struct {
	gen *GeneratorObj
}

type GeneratorObj struct {
	name     string
	filename string
	buf      bytes.Buffer

	Name GeneratorNameObj
}

func Init(name string, file string) *GeneratorObj {
	obj := GeneratorObj{}

	obj.name = name
	obj.filename = file

	obj.Name.gen = &obj

	return &obj
}
