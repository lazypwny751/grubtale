package imagination

import (
	"os"
	"fmt"
	"image"

	"math/rand"
	"image/png"
	// "image/draw"
	"image/color"

	// "golang.org/x/image/font"
	// "golang.org/x/image/font/basicfont"
	// "golang.org/x/image/math/fixed"
)

type BackgroundConfig struct {
	X_size      int
	Y_size      int
	FontFile    string
	TextColor   color.Color
	BgColor     color.Color
	BgImages    []string
}

type UserConfig struct {
	X_pos      int
	Y_pos      int
	FontSize   int
	ImagePath  string

	UserTitle  string      // Gece
	Version    string      // 1.0.0
	Hp         string      // 20/20
	Pkg        int         // 2386
}

type StatConfig struct {
	X_pos      int
	Y_pos      int
	FontSize   int
	ImagePath  string

	OsName     string      // Mint 20.3
	Cpu        string      // intel i3-10100
	Memory     string      // 8 GB
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

func Generator(out string, bg_config BackgroundConfig, user_config UserConfig, stat_config StatConfig) (error) {
	// Random background selection
	bg_img := bg_config.BgImages[rand.Intn(len(bg_config.BgImages))]
	if _, err := os.Stat(bg_img); os.IsNotExist(err) {
		return fmt.Errorf("background image %s does not exist", bg_img)
	}

	return nil
}