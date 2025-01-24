package i18n

import (
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path"
	"strings"

	_ "embed"
)

const (
	defaultI18nDir            = "i18n"
	defaultI18nOutRelativeDir = "i18n"
)

var (
	//go:embed en.toml
	en string
	//go:embed zh.toml
	zh string
)

func Generate(gen *protogen.Plugin) string {
	i18nDir, i18nOutRelativeDir := getI18nDir(gen)
	if err := os.MkdirAll(i18nDir, 0755); err != nil {
		gen.Error(err)
		return ""
	}

	if err := generateTOML(gen, i18nOutRelativeDir, "en", en); err != nil {
		gen.Error(err)
		return ""
	}

	if err := generateTOML(gen, i18nOutRelativeDir, "zh", zh); err != nil {
		gen.Error(err)
		return ""
	}

	return i18nDir
}

func getI18nDir(gen *protogen.Plugin) (string, string) {
	parts := strings.Split(gen.Request.GetParameter(), ",")

	dir, outDir := defaultI18nDir, defaultI18nOutRelativeDir
	for _, part := range parts {
		if strings.HasPrefix(part, "i18n_dir=") {
			dir = strings.TrimPrefix(part, "i18n_dir=")
		}
		if strings.HasPrefix(part, "i18n_out_relative_dir=") {
			outDir = strings.TrimPrefix(part, "i18n_out_relative_dir=")
		}
	}

	return dir, outDir
}

func generateTOML(gen *protogen.Plugin, dir, lang, content string) error {
	file := gen.NewGeneratedFile(path.Join(dir, lang+".toml"), "")
	file.P(content)
	return nil
}
