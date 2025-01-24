var (
	i18nBundle = newI18n("{{.I18nDir}}")
)

func newI18n(dir string) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
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
	return bundle
}
