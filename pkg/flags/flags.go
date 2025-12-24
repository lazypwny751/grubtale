package flags

import "flag"

var (
	Output = flag.String(
		"output",
		"/boot/grub/themes/Grubtale",
		"output directory name",
	)

	Title = flag.String(
		"title",
		"Grubtale",
		"theme title",
	)

	Config = flag.String(
		"config",
		"",
		"path to configuration file (json)",
	)
)

func Parse() {
	flag.Parse()
}
