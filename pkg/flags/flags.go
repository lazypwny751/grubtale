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

	Timeout = flag.Int(
		"timeout",
		-1,
		"timeout duration (seconds)",
	)

	Install = flag.Bool(
		"install",
		false,
		"install grubtale to system",
	)

	GrubPath = flag.String(
		"grub-path",
		"",
		"path to grub directory (default: auto-detect)",
	)

	InitSystem = flag.String(
		"init-system",
		"auto",
		"init system type (systemd, sysvinit, auto)",
	)
)

func Parse() {
	flag.Parse()
}
