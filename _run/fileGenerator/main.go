package main

import "github.com/hard-soft-ware/mpost/generator"

////////////////////////////////////

func enumGenerator(obj *generator.GeneratorObj, val []string) {
	obj.Print("type ").Name.Type().PrintLN(" byte").LN()

	obj.PrintLN("const (")
	for pos, text := range val {
		obj.Offset(1).Name.SelfCode(text).Print(" ")
		obj.Name.Type().Print(" = ").Number(pos).LN()
	}

	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")

	for _, text := range val {
		obj.Offset(1).Name.TextCode(text).Print(" = ")
		obj.String(text).LN()
	}

	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Name.Map().Print(" = map[").Name.Type().PrintLN("]string{")

	for _, text := range val {
		obj.Offset(1).Name.SelfCode(text).Print(": ")
		obj.Name.TextCode(text).PrintLN(",")
	}

	obj.PrintLN("}").LN()

	//

	obj.Print("func (obj ").Name.Type().PrintLN(") String() string {")
	obj.Offset(1).Print("val, ok := ").Name.Map().PrintLN("[obj]")
	obj.Offset(1).PrintLN("if ok {").Offset(2).PrintLN("return val").Offset(1).PrintLN("}")
	obj.Offset(1).Print("return \"Unknown ").Name.Type().PrintLN("\"")
	obj.PrintLN("}")

}

func constGenerator(obj *generator.GeneratorObj, val *generator.GeneratorValueObj) {
	obj.Print("type ").Name.Type().PrintLN(" int").LN()

	obj.PrintLN("const (")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(" ")
		obj.Name.Type().Print(" = ").Hex(code).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("const (")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.TextCode(text).Print(" = ")
		obj.String(text).LN()

		if val.Get.IsDelim(code) {
			obj.LN()
		}
	}
	obj.PrintLN(")").LN()

	//

	obj.Print("var ").Name.Map().Print(" = map[").Name.Type().PrintLN("]string{")
	for _, code := range val.Get.Ints() {
		text := val.Get.Text(code)

		obj.Offset(1).Name.SelfCode(text).Print(": ")
		obj.Name.TextCode(text).PrintLN(",")

	}
	obj.PrintLN("}").LN()

	obj.Print("func (obj ").Name.Type().PrintLN(") String() string {")
	obj.Offset(1).Print("val, ok := ").Name.Map().PrintLN("[obj]")
	obj.Offset(1).PrintLN("if ok {").Offset(2).PrintLN("return val").Offset(1).PrintLN("}")
	obj.Offset(1).Print("return \"Unknown ").Name.Type().PrintLN("\"")
	obj.PrintLN("}")

}

////////

func main() {
	Status()
	Document()
	Orientation()
	OrientationControl()
	PowerUp()
	PupExt()
	State()
	Bezel()
	Event()

	Cmd()
	CmdAux()
}
