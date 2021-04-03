package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Domain struct {
	Entities []Entity `yaml:"entities"`
}

const (
	genFilePath = "../play/User.yaml"
	destFilePath = "../target"
	instructionSplitToken = "\n"
)

var instructions []string
var currentInstruction int = 0

func main() {
	fmt.Print("Node js project generator")

	genFileData, err := ioutil.ReadFile(genFilePath)

	if err!=nil {
		print(err.Error())
		panic(0)
	}

	var domain Domain

	err = yaml.Unmarshal(genFileData, &domain)
	if err != nil {
		panic(err)
	}
	println()
	fmt.Printf("Value: %#v\n", domain.Entities)

	for _,v := range domain.Entities {
		v.PrintToFile()
	}

}