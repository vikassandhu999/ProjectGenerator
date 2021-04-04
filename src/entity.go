package main

import (
	"fmt"
	"io/ioutil"
)

type Entity struct {
	Name string    `yaml:"name"'`
	Fields []Field `yaml:"fields"`
	HasMapper bool `yaml:"has_mapper"`
}


func (e Entity) CreateMapper() {
	if e.HasMapper==false {
		return
	}
	mapperName:=fmt.Sprintf("%sMapper",e.Name)
	filename:=fmt.Sprintf("%s.ts",mapperName)

	domainMapper:=""
	persistenceMapper:=""

	for _, field := range e.Fields {
		domainMapper+=field.ToDomainMapperField()
		persistenceMapper+=field.ToPersistenceMapperField()
	}

	outputFileData:=fmt.Sprintf("export default class %s {\n" +
		"\tpublic static toDomain(model : any) {\n" +
		"\t\treturn new %s({\n%s\n})" +
		"\t}\n\n" +
		"\tpublic static toPersistence(domainModel : %s) {\n" +
		"\t\treturn {\n%s\n}" +
		"\t}\n\n" +
		"}\n",
		mapperName,
		e.Name,
		domainMapper,
		e.Name,
		persistenceMapper,
		)

	println()
	err := ioutil.WriteFile(destPath+"/mappers/"+filename,[]byte(outputFileData),0644)
	if err!=nil {
		print(err)
	}
}

func (e Entity) CreateDomain()  {
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
	err := ioutil.WriteFile(destPath+"/domain/"+filename,[]byte(outputFileData),0644)
	if err!=nil {
		print(err.Error())
	}
}
