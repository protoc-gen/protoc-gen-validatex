package i18n

import (
	"bufio"
	"bytes"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path"
	"sort"
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
	i18nDir, _ := getI18nDir(gen)
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
	tomlPath := path.Join(dir, lang+".toml")

	existingEntries, err := loadExistingTOML(tomlPath)
	if err != nil {
		gen.Error(err)
		return err
	}

	newEntries := parseTOMLEntries(content)

	for key, val := range newEntries {
		if _, exists := existingEntries[key]; !exists {
			existingEntries[key] = val // 只添加新 key
		}
	}

	keys := make([]string, 0, len(existingEntries))
	for key := range existingEntries {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	for _, key := range keys {
		buffer.WriteString(fmt.Sprintf("[%s]\nother = \"%s\"\n\n", key, existingEntries[key]))
	}

	if err := os.WriteFile(tomlPath, buffer.Bytes(), 0644); err != nil {
		gen.Error(err)
		return err
	}

	return nil
}

// loadExistingTOML parses an existing TOML file into a map of keys with their values.
func loadExistingTOML(filePath string) (map[string]string, error) {
	entries := make(map[string]string)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return entries, nil // File does not exist, return empty map
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open TOML file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentKey string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentKey = line[1 : len(line)-1]
		} else if strings.HasPrefix(line, "other = ") {
			if currentKey != "" {
				entries[currentKey] = strings.Trim(line[len("other = "):], "\"")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read TOML file: %w", err)
	}

	return entries, nil
}

func parseTOMLEntries(content string) map[string]string {
	entries := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(content))

	var currentKey string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentKey = line[1 : len(line)-1]
			entries[currentKey] = ""
		} else if strings.HasPrefix(line, "other = ") && currentKey != "" {
			entries[currentKey] = strings.Trim(line[len("other = "):], "\"")
		}
	}

	return entries
}
