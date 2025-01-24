package version

import (
	_ "embed"
)

var (
	//go:embed VERSION
	Version string
)
