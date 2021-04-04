package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Domain struct {
	Entities []Entity `yaml:"entities"`
}

var (
	genFilePath           = "../play/User.yaml"
	destPath              = "../target"
	instructionSplitToken = "\n"
)

func main() {

	args := os.Args[1:]
	if len(args)<2 {
		print("Invalid arguments")
		panic(0)
	}
	genFilePath=args[0]
	destPath=args[1]

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
		v.CreateDomain()
		v.CreateMapper()
	}

}