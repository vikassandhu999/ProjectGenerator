package main

import (
	"fmt"
	"io/ioutil"
)

type Entity struct {
	Name string    `yaml:"name"'`
	Fields []Field `yaml:"fields"`
}

func (e Entity) PrintToFile()  {
	filename:=fmt.Sprintf("%s.ts",e.Name)
	className:=fmt.Sprintf("%s",e.Name)
	interfaceName:=fmt.Sprintf("I%s",e.Name)
	propsName:=fmt.Sprintf("%sProps",e.Name)

	fieldDefString:=""
	gettersString:=""
	settersString:=""

	for _, field := range e.Fields {
		gettersString+=field.ToClassGetter()
		settersString+=field.ToClassSetter()
		fieldDefString+=field.ToInterfaceField()
	}

	entityInterface:=fmt.Sprintf("export interface %s {\n" +
		"%s\n" +
		"}",interfaceName,fieldDefString)

	entityProps:=fmt.Sprintf("interface %s {\n" +
		"%s\n" +
		"}",propsName,fieldDefString)

	outputFileData:=fmt.Sprintf("%s\n\n%s\n\nexport default class %s {\n" +
		"\tstate : %s;\n" +
		"\tconstructor(props : %s) {\n" +
		"\t\tthis.state = {...props};\n" +
		"\t}\n" +
		"\n\n%s\n" +
		"\n\n%s\n" +
		"\ttoDTO() : %s {\n" +
		"\t\treturn this.state;\n" +
		"\t}\n" +
		"}\n",
		entityInterface,
		entityProps,
		className,
		interfaceName,
		interfaceName,
		settersString,
		gettersString,
		interfaceName)

	println()
	err := ioutil.WriteFile(destFilePath+"/"+filename,[]byte(outputFileData),0644)
	if err!=nil {
		print(err)
	}
}
