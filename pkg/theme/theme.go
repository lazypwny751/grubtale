package theme

import "fmt"

// Theme configuration structures.
type GeneralThemeConfig struct {
	Title     string `json:"title"`
	CountDown int    `json:"count_down"`
	BgFile    string `json:"bg_file"`
	FontSize  int    `json:"font_size"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type BootThemeConfig struct {
	Top      int `json:"top"`
	Left     int `json:"left"`
	Width    int `json:"width"`
	Height   int `json:"height"`
	FontSize int `json:"font_size"`
}

type TimeoutThemeConfig struct {
	Duration int `json:"duration"`
	FontSize int `json:"font_size"`
}

type GrubtaleConfig struct {
	General GeneralThemeConfig `json:"general"`
	Boot    BootThemeConfig    `json:"boot"`
	Timeout TimeoutThemeConfig `json:"timeout"`
}

// GenerateTheme generates the theme configuration based on the provided general configuration.
func GenerateTheme(generalConf GeneralThemeConfig, bootConf BootThemeConfig, timeoutConf TimeoutThemeConfig) string {
	theme := general(generalConf)
	theme += "\n"
	theme += bootMenu(bootConf)
	theme += "\n"
	theme += timeoutSection(timeoutConf)

	return theme
}

// Individual theme section generators.
func general(config GeneralThemeConfig) string {
	// title-text
	theme := fmt.Sprintf(
		"title-text:    \"%s\"\n",
		config.Title,
	)

	// title-font
	theme += fmt.Sprintf(
		"title-font:    \"Determination Mono Regular %d\"\n",
		config.FontSize,
	)

	// title-color
	theme += fmt.Sprintf(
		"title-color:   \"#ffffffff\"\n",
	)

	// desktop-color
	theme += fmt.Sprintf(
		"desktop-color: \"#000000ff\"\n",
	)

	// desktop-image
	theme += fmt.Sprintf(
		"desktop-image: \"background.png\"\n",
	)

	return theme
}

func bootMenu(config BootThemeConfig) string {
	// boot_menu
	theme := fmt.Sprintf(
		"+ boot_menu {\n",
	)

	// =* boot_menu properties *=//
	// top
	theme += fmt.Sprintf(
		"   top                 = %d%s\n",
		config.Top,
		"%",
	)

	// left
	theme += fmt.Sprintf(
		"   left                = %d%s\n",
		config.Left,
		"%",
	)

	// height
	theme += fmt.Sprintf(
		"   height              = %d%s\n",
		config.Height,
		"%",
	)

	// width
	theme += fmt.Sprintf(
		"   width               = %d%s\n",
		config.Width,
		"%",
	)

	// menu_pixmap_style
	theme += fmt.Sprintf(
		"   menu_pixmap_style   = \"menu_*.png\"\n",
	)

	// item_font
	theme += fmt.Sprintf(
		"   item_font           = \"Determination Mono Regular %d\"\n",
		config.FontSize,
	)

	// item_color
	theme += fmt.Sprintf(
		"   item_color          = \"#ffffffff\"\n",
	)

	// selected_item_color
	theme += fmt.Sprintf(
		"   selected_item_color = \"#808080\"\n",
	)

	// item_height
	theme += fmt.Sprintf(
		"   item_height         = 28\n",
	)

	// item_padding
	theme += fmt.Sprintf(
		"   item_padding        = 16\n",
	)

	// item_spacing
	theme += fmt.Sprintf(
		"   item_spacing        = 6\n",
	)

	// end of boot_menu
	theme += fmt.Sprintf(
		"}\n",
	)

	return theme
}

func timeoutSection(config TimeoutThemeConfig) string {
	// timeout label
	theme := fmt.Sprintf(
		"+ label {\n",
	)

	// =* timeout label properties *=//
	// top
	theme += fmt.Sprintf(
		"   top   = %s\n",
		"1%",
	)

	// left
	theme += fmt.Sprintf(
		"   left  = %s\n",
		"1%",
	)

	// timeout label id
	theme += fmt.Sprintf(
		"   id    = \"__timeout__\"\n",
	)

	// timeout label text
	theme += fmt.Sprintf(
		"   text  = \"QUITTING.. %s/%d\"\n",
		"%d",
		config.Duration,
	)

	// timeout label style
	theme += fmt.Sprintf(
		"   align = \"right\"\n",
	)

	// timeout label font
	theme += fmt.Sprintf(
		"   font  = \"Determination Mono Regular %d\"\n",
		config.FontSize,
	)

	// timeout label color
	theme += fmt.Sprintf(
		"   color = \"#ffffffff\"\n",
	)

	// end of timeout label
	theme += fmt.Sprintf(
		"}\n",
	)

	return theme
}
