package validatex

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path"
)

var (
	bundle *i18n.Bundle
)

func Init18n(dir string) {
	if bundle != nil {
		return
	}
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if path.Ext(file.Name()) != ".toml" {
			continue
		}
		bundle.MustLoadMessageFile(path.Join(dir, file.Name()))
	}
}

func GetBundle() *i18n.Bundle {
	return bundle
}

func Localize(ctx context.Context, c *i18n.LocalizeConfig) (string, error) {
	return i18n.NewLocalizer(bundle,
		GetValFromCtx(ctx, KeyXLang), GetValFromCtx(ctx, KeyAcceptLang)).
		Localize(c)
}

func MustLocalize(ctx context.Context, c *i18n.LocalizeConfig, defVal string) string {
	s, err := Localize(ctx, c)
	if err != nil {
		return defVal
	}
	return s
}
