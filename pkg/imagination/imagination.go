package imagination

import (
	"os"
	"image"
	"image/png"
	"image/draw"
	"image/color"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type UserConfig struct {
	X_pos      int
	Y_pos      int
	FontSize   int
	ImagePath  string

	UserTitle  string      // Gece
	Version    int         // 1.0.0
	Hp         (int, int)  // 20/20
	Pkg        int         // 2386
}

type StatConfig struct {
	X_pos      int
	Y_pos      int
	FontSize   int
	ImagePath  map[string]string

	OsName     string      // Mint 20.3
	Cpu        string      // intel i3-10100
	Memory     string      // 8 GB
}

type ImageConfig struct {
	X_size      int
	Y_size      int
	FontFile    string
	TextColor   color.Color
	BgColor     color.Color
	BgImages    map[string]image.Image
}

// Public helpers for image manipulation.
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func DrawText(img draw.Image, text string, x, y int, col color.Color) {
	point := fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}

func Generator(...) (error) {
	return nil
}