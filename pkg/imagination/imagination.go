package imagination

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"image/color"
	"image/draw"
	"image/png"
	"math/rand"

	"github.com/golang/freetype/truetype"
	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type BackgroundConfig struct {
	FontFile  []byte
	TextColor color.Color
	BgColor   color.Color
	BgImages  [][]byte
}

type UserConfig struct {
	FontSize  int
	ImagePath []byte

	UserTitle string // Gece
	Version   string // 1.0.0
	Pkg       int    // 4442
}

type StatConfig struct {
	FontSize  int
	ImagePath []byte

	OsName string // Mint 20.3
	Cpu    string // intel i3-10100
	Memory string // 8 GB
}

// Public helpers for image manipulation.
func LoadImageFromBytes(data []byte) (image.Image, error) {
	img, err := png.Decode(bytes.NewReader(data))
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

func Generator(out string, bg_config BackgroundConfig, user_config UserConfig, stat_config StatConfig) error {
	// Random background selection
	bg_data := bg_config.BgImages[rand.Intn(len(bg_config.BgImages))]

	bg_img, err := LoadImageFromBytes(bg_data)
	if err != nil {
		return fmt.Errorf("failed to decode background image: %v", err)
	}

	// Create a mutable image
	bounds := bg_img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, bg_img, image.Point{}, draw.Src)

	// Load Font
	f, err := truetype.Parse(bg_config.FontFile)
	if err != nil {
		return fmt.Errorf("failed to parse font: %v", err)
	}

	// Calculate responsive positions and sizes
	// Target width for boxes: 25% of screen width (Smaller)
	targetBoxWidth := float64(width) * 0.25

	// Margins
	marginX := int(float64(width) * 0.05)    // 5% margin from left
	marginTop := int(float64(height) * 0.10) // 10% margin from top (Aligned with menu)

	// User Box Position
	userX := marginX
	userY := marginTop
	var userHeight int

	// Draw User Config (User Box)
	if len(user_config.ImagePath) > 0 {
		userImg, err := LoadImageFromBytes(user_config.ImagePath)
		if err == nil {
			// Scale user image
			scale := targetBoxWidth / float64(userImg.Bounds().Dx())
			newWidth := int(targetBoxWidth)
			newHeight := int(float64(userImg.Bounds().Dy()) * scale)
			userHeight = newHeight

			scaledUserImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
			xdraw.CatmullRom.Scale(scaledUserImg, scaledUserImg.Bounds(), userImg, userImg.Bounds(), xdraw.Over, nil)

			// Draw user image
			userRect := image.Rect(userX, userY, userX+newWidth, userY+newHeight)
			draw.Draw(rgba, userRect, scaledUserImg, image.Point{}, draw.Over)

			// Scale font size relative to box scaling
			// Base font size was 24 for 800x600.
			// Let's make it relative to box height or width.
			// Or just use the provided FontSize but scaled?
			// User provided FontSize 24. Let's scale it by the same factor we scaled the image?
			// Assuming original image was designed for 1:1 scale.
			scaledFontSize := float64(user_config.FontSize) * scale * 0.8 // Slightly smaller to fit

			face := truetype.NewFace(f, &truetype.Options{
				Size:    scaledFontSize,
				DPI:     72,
				Hinting: font.HintingFull,
			})

			d := &font.Drawer{
				Dst:  rgba,
				Src:  image.NewUniform(bg_config.TextColor),
				Face: face,
			}

			// Adjust text positions relative to the box
			paddingX := int(float64(newWidth) * 0.1)
			paddingY := int(float64(newHeight) * 0.25) // Start text a bit lower
			lineHeight := int(scaledFontSize * 1.5)

			// Name (No quotes)
			d.Dot = fixed.P(userX+paddingX, userY+paddingY)
			d.DrawString(user_config.UserTitle)

			// VER (Version Instead of Level"LV")
			d.Dot = fixed.P(userX+paddingX, userY+paddingY+lineHeight)
			d.DrawString(fmt.Sprintf("VER %s", user_config.Version))

			// PKG (Instead of HP)
			d.Dot = fixed.P(userX+paddingX, userY+paddingY+lineHeight*2)
			d.DrawString(fmt.Sprintf("PKG %d", user_config.Pkg))
		}
	}

	// Stat Box Position: Below User Box with some gap
	statY := userY + userHeight + int(float64(height)*0.02) // 2% gap (Closer)

	// Draw Stat Config (Stat Box)
	if len(stat_config.ImagePath) > 0 {
		statImg, err := LoadImageFromBytes(stat_config.ImagePath)
		if err == nil {
			// Scale stat image
			scale := targetBoxWidth / float64(statImg.Bounds().Dx())
			newWidth := int(targetBoxWidth)
			newHeight := int(float64(statImg.Bounds().Dy()) * scale)

			scaledStatImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
			xdraw.CatmullRom.Scale(scaledStatImg, scaledStatImg.Bounds(), statImg, statImg.Bounds(), xdraw.Over, nil)

			// Align stat box with user box
			statX := marginX

			// Draw stat image
			statRect := image.Rect(statX, statY, statX+newWidth, statY+newHeight)
			draw.Draw(rgba, statRect, scaledStatImg, image.Point{}, draw.Over)

			// Larger font for stat box content
			scaledFontSize := float64(stat_config.FontSize) * scale * 1.1

			// Draw text on stat image
			face := truetype.NewFace(f, &truetype.Options{
				Size:    scaledFontSize,
				DPI:     72,
				Hinting: font.HintingFull,
			})

			d := &font.Drawer{
				Dst:  rgba,
				Src:  image.NewUniform(bg_config.TextColor),
				Face: face,
			}

			paddingX := int(float64(newWidth) * 0.1)
			paddingY := int(float64(newHeight) * 0.25)
			lineHeight := int(scaledFontSize * 1.5)

			// OS Name
			d.Dot = fixed.P(statX+paddingX, statY+paddingY)
			d.DrawString(stat_config.OsName)

			// CPU
			d.Dot = fixed.P(statX+paddingX, statY+paddingY+lineHeight)
			d.DrawString(stat_config.Cpu)

			// Memory / GRUBTALE
			d.Dot = fixed.P(statX+paddingX, statY+paddingY+lineHeight*2)
			d.DrawString(stat_config.Memory)
		}
	}

	return SaveImage(rgba, out)
}
