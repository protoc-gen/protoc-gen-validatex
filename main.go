package main

import (
	"github.com/protoc-gen/protoc-gen-validatex/i18n"
	"github.com/protoc-gen/protoc-gen-validatex/validatex"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
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
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
