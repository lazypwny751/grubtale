package main

// Grubtale an undertale inspired grub theme.
// This theme generator aims to bring the spirit of Undertale into the GRUB bootloader.

import (
	"os"
	"fmt"

	"log/slog"
	"image/color"
	"path/filepath"

	"github.com/lazypwny751/grubtale/pkg/theme"
	"github.com/lazypwny751/grubtale/pkg/flags"
	"github.com/lazypwny751/grubtale/pkg/imagination"
)


func main() {
	// Parse command line flags.
	flags.Parse()

	if err := os.MkdirAll(*flags.Output, os.ModePerm); err != nil {
		slog.Error("dir", fmt.Sprintf("Could not create directory %s\n", *flags.Output), flags.Output)
		return
	}

	stat, err := os.Stat(*flags.Output)
	if err != nil {
		slog.Error("dir", fmt.Sprintf("Could not access directory %s\n", *flags.Output), flags.Output)
		return
	}

	if !stat.IsDir() {
		slog.Error("dir", fmt.Sprintf("%s is not a directory\n", *flags.Output), flags.Output)
		return
	}

	// =* Generate background file. *=//
	backgroundConfig := imagination.BackgroundConfig{
		X_size:    800,
		Y_size:    600,
		FontFile:  "assets/ttf/determination-mono/determination-mono.ttf",
		TextColor: color.RGBA{255, 255, 255, 255},
		BgColor:   nil,
		BgImages:  []string{
			"assets/background/before-muffet.png",
			"assets/background/muffet-home-door.png",
		},
	}

	userConfig := imagination.UserConfig{
		X_pos:     50,
		Y_pos:     50,
		FontSize:  24,
		ImagePath: "assets/png/user.png",

		UserTitle: "Gece",
		Version:   "1.0.0",
		Hp:        "20/20",
		Pkg:       2386,
	}

	statConfig := imagination.StatConfig{
		X_pos:     50,
		Y_pos:     150,
		FontSize:  18,
		ImagePath: "assets/png/stat.png",

		OsName:  "Mint 20.3",
		Cpu:     "intel i3-10100",
		Memory:  "8 GB",
	}

	if err := imagination.Generator(filepath.Join(*flags.Output, "background.png"), backgroundConfig, userConfig, statConfig); err != nil {
		slog.Error("img", "Could not generate background image", "error", err)
		return
	}

	// =* Generate theme configuration. *=//
	// general theme config
	generalThemeConfig := theme.GeneralThemeConfig{
		Title:      "", // *flags.Title
		CountDown:  20,
		BgFile:    "background.png",
		FontSize:  24,
	}

	bootThemeConfig := theme.BootThemeConfig{
		Top:      10,
		Left:     25,
		FontSize: 24,
	}

	timeoutThemeConfig := theme.TimeoutThemeConfig{
		Duration:  30,
		FontSize:  18,
	}

	// =* Generate theme data. *=//
	themeData := theme.GenerateTheme(generalThemeConfig, bootThemeConfig, timeoutThemeConfig)
	fmt.Printf(themeData)
}