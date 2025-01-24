package main

import (
	"github.com/protoc-gen/protoc-gen-validatex/i18n"
	"github.com/protoc-gen/protoc-gen-validatex/validatex"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		i18nDir := i18n.Generate(gen)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			validatex.GenerateFile(gen, f, i18nDir)
		}
		return nil
	})
}
