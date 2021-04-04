package main

import "fmt"

type Field struct {
	DomainId string `yaml:"domain_id"`
	MapperId string `yaml:"mapper_id"`
	Type string `yaml:"type"`
	AccessModifier string `yaml:"access_modifier"`
	Optional bool `yaml:"optional"`
}

func (f Field) ToInterfaceField() string {
	var sign string = ""
	if f.Optional==true {
		sign="?"
	}
	return fmt.Sprintf("\t%s %s: %s;\n",f.DomainId,sign,f.Type)
}

func (f Field) ToClassField() string {
	return fmt.Sprintf("\t%s %s : %s;\n",f.AccessModifier,f.DomainId,f.Type)
}

func (f Field) ToClassGetter() string {
	return fmt.Sprintf("\tget %s(): %s {\n" +
		"\t\treturn this.state.%s;\n" +
		"\t}\n",f.DomainId,f.Type,f.DomainId)
}

func (f Field) ToClassSetter() string {
	return fmt.Sprintf("\tset %s(%s:%s) {\n" +
		"\t\tthis.state.%s=%s;\n" +
		"\t}\n",f.DomainId,f.DomainId,f.Type,f.DomainId,f.DomainId)
}

func (f Field) ToDomainMapperField() string {
	return fmt.Sprintf("\t\t%s: model.%s,\n",f.DomainId,f.MapperId)
}

func (f Field) ToPersistenceMapperField() string {
	return fmt.Sprintf("\t\t%s: domainModel.%s,\n",f.MapperId,f.DomainId)
}

func InstructionToField() Field {
	return Field{}
}
