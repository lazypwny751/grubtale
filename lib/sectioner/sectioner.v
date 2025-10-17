module sectioner

import toml

struct Theme {
	font string
}

struct BootMenu {
	top    string
	left   string
	height string
	width  string
}

struct User {
	top	    string
	left    string
	height  string
	width   string
	user    map[string]string
	version map[string]string
	hp      map[string]string
	pkg     map[string]string
}

struct State {
	top	             string
	left             string
	height           string
	width            string
	operating_system map[string]string
	cpu              map[string]string
	memory           map[string]string
}

struct General {
	background_image string
	theme            Theme
}

pub struct Config {
	general  General
	bootmenu BootMenu
	user     User
	state    State
}

pub fn parse_config(cfg toml.Doc) !Config {
	mut config := Config{
		general:  General{}
		bootmenu: BootMenu{}
		user:     User{}
		state:    State{}
	}
	return config
}