package i18n

import (
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path"
	"strings"

	_ "embed"
)

const (
	defaultI18nDir = "i18n"
)

var (
	//go:embed en.toml
	en string
	//go:embed zh.toml
	zh string
)

func Generate(gen *protogen.Plugin) string {
	i18nDir := getI18nDir(gen)
	if err := os.MkdirAll(i18nDir, 0755); err != nil {
		gen.Error(err)
		return ""
	}

	if err := generateTOML(gen, i18nDir, "en", en); err != nil {
		gen.Error(err)
		return ""
	}

	if err := generateTOML(gen, i18nDir, "zh", zh); err != nil {
		gen.Error(err)
		return ""
	}

	return i18nDir
}

func getI18nDir(gen *protogen.Plugin) string {
	parts := strings.Split(gen.Request.GetParameter(), ",")

	for _, part := range parts {
		if strings.HasPrefix(part, "i18n_dir=") {
			return strings.TrimPrefix(part, "i18n_dir=")
		}
	}

	return defaultI18nDir
}

func generateTOML(gen *protogen.Plugin, dir, lang, content string) error {
	file := gen.NewGeneratedFile(path.Join(dir, lang+".toml"), "")
	file.P(content)
	return nil
}
