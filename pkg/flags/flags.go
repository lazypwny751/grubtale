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
)

func Parse() {
	flag.Parse()
}