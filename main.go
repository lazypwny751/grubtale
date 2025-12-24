package main

// Grubtale an undertale inspired grub theme.
// This theme generator aims to bring the spirit of Undertale into the GRUB bootloader.

import (
	"encoding/json"
	"os"

	"image/color"
	"log/slog"
	"path/filepath"

	"github.com/lazypwny751/grubtale/assets"
	"github.com/lazypwny751/grubtale/pkg/flags"
	"github.com/lazypwny751/grubtale/pkg/generator"
	"github.com/lazypwny751/grubtale/pkg/imagination"
	"github.com/lazypwny751/grubtale/pkg/theme"
)

func main() {
	// Parse command line flags.
	flags.Parse()

	if err := os.MkdirAll(*flags.Output, os.ModePerm); err != nil {
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
	bg1, err := assets.ReadFile("background/before-muffet.png")
	if err != nil {
		slog.Error("Could not read bg asset", "error", err)
		return
	}
	bg2, err := assets.ReadFile("background/muffet-home-door.png")
	if err != nil {
		slog.Error("Could not read bg asset", "error", err)
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
	// Default configuration
	grubtaleConfig := theme.GrubtaleConfig{
		General: theme.GeneralThemeConfig{
			Title:     "", // *flags.Title
			CountDown: 20,
			BgFile:    "background.png",
			FontSize:  24,
		},
		Boot: theme.BootThemeConfig{
			Top:      10,
			Left:     35,
			Width:    60,
			Height:   80,
			FontSize: 24,
		},
		Timeout: theme.TimeoutThemeConfig{
			Duration: 30,
			FontSize: 16,
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

	// =* Generate background file. *=//
	backgroundConfig := imagination.BackgroundConfig{
		FontFile:  fontData,
		TextColor: color.RGBA{255, 255, 255, 255},
		BgColor:   nil,
		BgImages:  [][]byte{bg1, bg2},
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
		FontSize:  18,
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
	menuFiles := []string{
		"menu_c.png", "menu_e.png", "menu_n.png", "menu_ne.png",
		"menu_nw.png", "menu_s.png", "menu_se.png", "menu_sw.png", "menu_w.png",
	}

	// Calculate scale factor for menu images
	bgWidth := grubtaleConfig.General.Width
	if bgWidth == 0 {
		// If not specified, use the original background width
		img, err := imagination.LoadImageFromBytes(bg1)
		if err == nil {
			bgWidth = img.Bounds().Dx()
		} else {
			bgWidth = 1920 // Fallback
		}
	}

	var scale float64 = 1.0
	if bgWidth > 0 {
		userImgDecoded, err := imagination.LoadImageFromBytes(userImg)
		if err == nil {
			targetBoxWidth := float64(bgWidth) * 0.25
			scale = targetBoxWidth / float64(userImgDecoded.Bounds().Dx())
		}
	}

	for _, file := range menuFiles {
		data, err := assets.ReadFile("png/" + file)
		if err != nil {
			slog.Error("Could not read menu asset", "file", file, "error", err)
			continue
		}

		// Resize menu image
		scaledImg, err := imagination.ScaleImage(data, scale)
		if err != nil {
			slog.Error("Could not scale menu asset", "file", file, "error", err)
			continue
		}

		if err := imagination.SaveImage(scaledImg, filepath.Join(*flags.Output, file)); err != nil {
			slog.Error("Could not write menu asset", "file", file, "error", err)
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
