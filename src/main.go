package main

// Grubtale an undertale inspired grub theme.
// This theme generator aims to bring the spirit of Undertale into the GRUB bootloader.

import (
	"os"
	"fmt"

	"log/slog"

	"github.com/lazypwny751/grubtale/pkg/flags"
)


func main() {
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
}