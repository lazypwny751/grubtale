package main

/*
 * Grubtale
 *
 * An Undertale inspired GRUB theme generator.
 *
 * Created by lazypwny751 & GitHub Copilot(my best friend).
 *
 * This project was peer-coded with AI assistance, combining human creativity
 * with machine efficiency. Not just generated, but crafted together.
 *
 * Stay determined!
 */

import (
	"embed"
	"encoding/json"
	"os"

	"image/color"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/lazypwny751/grubtale/assets"
	"github.com/lazypwny751/grubtale/pkg/flags"
	"github.com/lazypwny751/grubtale/pkg/generator"
	"github.com/lazypwny751/grubtale/pkg/imagination"
	"github.com/lazypwny751/grubtale/pkg/selfinstall"
	"github.com/lazypwny751/grubtale/pkg/theme"
)

//go:embed grubtale.service grubtale.sh
var serviceFiles embed.FS

func main() {
	// Parse command line flags.
	flags.Parse()

	// Handle installation
	if *flags.Install {
		if err := selfinstall.Install(serviceFiles, *flags.Config, *flags.GrubPath, *flags.InitSystem); err != nil {
			slog.Error("Installation failed", "error", err)
			os.Exit(1)
		}
		return
	}

	if err := os.MkdirAll(*flags.Output, 0755); err != nil {
		slog.Error("Could not create directory", "path", *flags.Output, "error", err)
		return
	}

	stat, err := os.Stat(*flags.Output)
	if err != nil {
		slog.Error("Could not access directory", "path", *flags.Output, "error", err)
		return
	}

	if !stat.IsDir() {
		slog.Error("Output path is not a directory", "path", *flags.Output)
		return
	}

	// Load assets
	fontData, err := assets.ReadFile("ttf/determination-mono/determination-mono.ttf")
	if err != nil {
		slog.Error("Could not read font asset", "error", err)
		return
	}

	// Load background images dynamically
	var bgImages [][]byte
	bgFiles, err := assets.ReadDir("background")
	if err != nil {
		slog.Error("Could not read background directory", "error", err)
		return
	}

	for _, file := range bgFiles {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".png") {
			continue
		}
		data, err := assets.ReadFile("background/" + file.Name())
		if err != nil {
			slog.Error("Could not read bg asset", "file", file.Name(), "error", err)
			continue
		}
		bgImages = append(bgImages, data)
	}

	if len(bgImages) == 0 {
		slog.Error("No background images found in assets/background")
		return
	}

	userImg, err := assets.ReadFile("png/user.png")
	if err != nil {
		slog.Error("Could not read user asset", "error", err)
		return
	}
	statImg, err := assets.ReadFile("png/stat.png")
	if err != nil {
		slog.Error("Could not read stat asset", "error", err)
		return
	}

	// =* Generate theme configuration. *=//
	// Determine default timeout
	defaultTimeout := 30
	if grubTimeout := generator.GetGrubTimeout(); grubTimeout > 0 {
		defaultTimeout = grubTimeout
	}

	// Default configuration
	grubtaleConfig := theme.GrubtaleConfig{
		General: theme.GeneralThemeConfig{
			Title:     "",
			CountDown: 20,
			BgFile:    "background.png",
			FontSize:  32,
		},
		Boot: theme.BootThemeConfig{
			Top:      10,
			Left:     34,
			Width:    60,
			Height:   80,
			FontSize: 32,
		},
		Timeout: theme.TimeoutThemeConfig{
			Duration: defaultTimeout,
			FontSize: 24,
		},
	}

	// Try to load from config file if provided
	if *flags.Config != "" {
		if configFile, err := os.ReadFile(*flags.Config); err == nil {
			if err := json.Unmarshal(configFile, &grubtaleConfig); err != nil {
				slog.Warn("Failed to parse config file, using defaults", "path", *flags.Config, "error", err)
			} else {
				slog.Info("Loaded configuration", "path", *flags.Config)
			}
		} else {
			slog.Warn("Could not read config file, using defaults", "path", *flags.Config, "error", err)
		}
	}

	// Override with flag if provided
	if *flags.Timeout != -1 {
		grubtaleConfig.Timeout.Duration = *flags.Timeout
	}

	// Calculate scale factor
	bgWidth := grubtaleConfig.General.Width
	if bgWidth == 0 {
		// If not specified, use the original background width
		// Use the first loaded background image for reference
		img, err := imagination.LoadImageFromBytes(bgImages[0])
		if err == nil {
			bgWidth = img.Bounds().Dx()
		} else {
			bgWidth = 1920 // Default fallback width
		}
	}

	var scale float64 = 1.0
	if bgWidth > 0 {
		userImgDecoded, err := imagination.LoadImageFromBytes(userImg)
		if err == nil {
			targetBoxWidth := float64(bgWidth) * 0.20
			scale = targetBoxWidth / float64(userImgDecoded.Bounds().Dx())
		}
	}

	// Update Boot Config with scaled values
	if grubtaleConfig.Boot.ItemHeight == 0 {
		grubtaleConfig.Boot.ItemHeight = int(42 * scale)
	}
	if grubtaleConfig.Boot.ItemPadding == 0 {
		grubtaleConfig.Boot.ItemPadding = int(12 * scale)
	}
	if grubtaleConfig.Boot.ItemSpacing == 0 {
		grubtaleConfig.Boot.ItemSpacing = int(10 * scale)
	}
	// Scale font size if it's the default
	if grubtaleConfig.General.FontSize == 32 {
		grubtaleConfig.General.FontSize = getClosestFontSize(int(32 * scale))
	}
	if grubtaleConfig.Boot.FontSize == 32 {
		grubtaleConfig.Boot.FontSize = getClosestFontSize(int(32 * scale))
	}
	if grubtaleConfig.Timeout.FontSize == 24 {
		grubtaleConfig.Timeout.FontSize = getClosestFontSize(int(24 * scale))
	}

	// =* Generate background file. *=//
	backgroundConfig := imagination.BackgroundConfig{
		FontFile:  fontData,
		TextColor: color.RGBA{255, 255, 255, 255},
		BgColor:   nil,
		BgImages:  bgImages,
		Width:     grubtaleConfig.General.Width,
		Height:    grubtaleConfig.General.Height,
	}

	// Determine User Title
	userTitle := generator.GetHostname()
	if *flags.Title != "Grubtale" {
		userTitle = *flags.Title
	}

	userConfig := imagination.UserConfig{
		FontSize:  24,
		ImagePath: userImg,

		UserTitle: userTitle,
		Version:   "1.0.0",
		Pkg:       generator.GetPackageCount(),
	}

	statConfig := imagination.StatConfig{
		FontSize:  16,
		ImagePath: statImg,

		OsName: generator.GetOSName(),
		Cpu:    generator.GetCPUInfo(),
		Memory: "GRUBTALE",
	}

	if err := imagination.Generator(filepath.Join(*flags.Output, "background.png"), backgroundConfig, userConfig, statConfig); err != nil {
		slog.Error("Could not generate background image", "error", err)
		return
	}

	// =* Generate theme data. *=//
	themeData := theme.GenerateTheme(grubtaleConfig.General, grubtaleConfig.Boot, grubtaleConfig.Timeout)

	// Write theme.txt
	themePath := filepath.Join(*flags.Output, "theme.txt")
	if err := os.WriteFile(themePath, []byte(themeData), 0644); err != nil {
		slog.Error("Could not write theme.txt", "error", err)
		return
	}
	slog.Info("Generated theme.txt", "path", themePath)

	// Copy menu images
	pngFiles, err := assets.ReadDir("png")
	if err != nil {
		slog.Error("Could not read png directory", "error", err)
	} else {
		for _, fileEntry := range pngFiles {
			fileName := fileEntry.Name()
			if !strings.HasPrefix(fileName, "menu_") {
				continue
			}

			data, err := assets.ReadFile("png/" + fileName)
			if err != nil {
				slog.Error("Could not read menu asset", "file", fileName, "error", err)
				continue
			}

			// Resize menu image
			// Apply a multiplier to make them slightly larger as requested
			menuScale := scale * 1.5
			scaledImg, err := imagination.ScaleImage(data, menuScale)
			if err != nil {
				slog.Error("Could not scale menu asset", "file", fileName, "error", err)
				continue
			}

			if err := imagination.SaveImage(scaledImg, filepath.Join(*flags.Output, fileName)); err != nil {
				slog.Error("Could not write menu asset", "file", fileName, "error", err)
			}
		}
	}

	// Copy pf2 fonts
	pf2Files, err := assets.ReadDir("pf2")
	if err != nil {
		slog.Error("Could not read pf2 directory", "error", err)
	} else {
		for _, file := range pf2Files {
			if file.IsDir() {
				continue
			}
			data, err := assets.ReadFile("pf2/" + file.Name())
			if err != nil {
				slog.Error("Could not read pf2 asset", "file", file.Name(), "error", err)
				continue
			}
			if err := os.WriteFile(filepath.Join(*flags.Output, file.Name()), data, 0644); err != nil {
				slog.Error("Could not write pf2 asset", "file", file.Name(), "error", err)
			}
		}
	}
}

func getClosestFontSize(size int) int {
	available := []int{8, 16, 24, 32, 44, 64}
	closest := available[0]
	minDiff := abs(size - closest)
	for _, s := range available {
		diff := abs(size - s)
		if diff <= minDiff {
			minDiff = diff
			closest = s
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
