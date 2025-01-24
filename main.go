package main

import (
	"github.com/protoc-gen/protoc-gen-validatex/validatex"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			validatex.GenerateFile(gen, f)
		}
		return nil
	})
}
