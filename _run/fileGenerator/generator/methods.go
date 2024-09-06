package generator

import (
	"fmt"
	"strings"
)

//###########################################################//

func (gen *GeneratorObj) NewVal() *GeneratorValueObj {
	return generatorValueInit(&gen.name)
}

////

func (gen *GeneratorObj) Len() int {
	return gen.buf.Len()
}

func (gen *GeneratorObj) Write(data []byte) *GeneratorObj {
	gen.buf.Write(data)
	return gen
}

func (gen *GeneratorObj) Del(len int) *GeneratorObj {
	gen.buf.Truncate(gen.Len() - len)
	return gen
}

//

func (gen *GeneratorObj) LN() *GeneratorObj {
	gen.Write([]byte("\n"))
	return gen
}

func (gen *GeneratorObj) Print(text string) *GeneratorObj {
	gen.Write([]byte(text))
	return gen
}

func (gen *GeneratorObj) PrintLN(text string) *GeneratorObj {
	gen.Print(text).LN()
	return gen
}

func (gen *GeneratorObj) Repeat(chat string, size int) *GeneratorObj {
	gen.Print(strings.Repeat(chat, size))
	return gen
}

func (gen *GeneratorObj) Offset(size int) *GeneratorObj {
	gen.Repeat("\t", size)
	return gen
}

//

func (gen *GeneratorObj) String(a any) *GeneratorObj {
	gen.Print(fmt.Sprintf("\"%s\"", a))
	return gen
}

func (gen *GeneratorObj) Hex(a any) *GeneratorObj {
	gen.Print(fmt.Sprintf("0x%02x", a))
	return gen
}

func (gen *GeneratorObj) Number(a any) *GeneratorObj {
	gen.Print(fmt.Sprintf("%d", a))
	return gen
}
