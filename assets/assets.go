package assets

import (
	"embed"
	"io/fs"
)

//go:embed pf2/*.pf2
//go:embed png/*.png
//go:embed background/*.png
//go:embed ttf/determination-mono/*.ttf
//go:embed ttf/determination-mono/*.txt
var assets embed.FS

func ReadFile(name string) ([]byte, error) {
	return assets.ReadFile(name)
}

func ReadDir(name string) ([]fs.DirEntry, error) {
	return assets.ReadDir(name)
}
