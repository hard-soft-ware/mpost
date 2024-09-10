package generator

import "os"

//###########################################################//

type GeneratorSaveObj struct {
	gen *GeneratorObj
	Add GeneratorSaveAddObj

	pack    string
	imports []string
	types   []string
}

func (gen *GeneratorObj) Save(pack string) *GeneratorSaveObj {
	save := GeneratorSaveObj{}
	save.Add.save = &save

	save.gen = gen
	save.pack = pack

	return &save
}

func (save *GeneratorSaveObj) End() error {
	file, err := os.OpenFile(save.gen.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("package " + save.pack + "\n\n")

	switch {
	case len(save.imports) == 1:
		file.WriteString("import \"" + save.imports[0] + "\";\n\n")

	case len(save.imports) > 1:
		file.WriteString("import (\n")
		for _, i := range save.imports {
			file.WriteString("\t\"" + i + "\"\n")
		}
		file.WriteString(")\n\n")
	}

	file.WriteString(hText + "\n\n")

	if len(save.types) > 0 {
		for _, i := range save.types {
			file.WriteString("type " + i + "\n")
		}
		file.WriteString("\n")
	}

	file.Write(save.gen.buf.Bytes())
	return nil
}

//

type GeneratorSaveAddObj struct {
	save *GeneratorSaveObj
}

func (add *GeneratorSaveAddObj) Type(name string, types string) *GeneratorSaveObj {
	add.save.types = append(add.save.types, name+" "+types)
	return add.save
}

func (add *GeneratorSaveAddObj) Import(imports string) *GeneratorSaveObj {
	add.save.imports = append(add.save.imports, imports)
	return add.save
}
