module main

import os
import toml
import flag
import x.templating.dtm
import lib.osutils
import lib.sectioner

@[xdoc: "Undertale theme for GNU GRUB."]
@[footer: "\nStay determined!"]
@[version: "1.0.0"]
@[name: "Grubtale"]
struct Flags {
    show_help    bool   @[short: h; long: "help"; xdoc: "Show this help message and exit."]
    cli_mode     bool   @[short: c; long: "cli"; xdoc: "Run in CLI mode without graphical interface."]
    theme_path   string = '/boot/grub/themes/grubtale' @[short: d; long: "directory"; xdoc: "Path to the GRUB theme directory."]
    config_path  string = 'profile.toml' @[short: p; long: "config"; xdoc: "Path to the configuration file."]
}

fn main() {
    flags, no_matches := flag.to_struct[Flags](os.args, skip: 1)!

    if no_matches.len > 0 {
        println('The following flags could not be mapped to any fields on the struct: ${no_matches}')
    }

    if flags.show_help {
        documentation := flag.to_doc[Flags]()!
        println(documentation)
        exit(0)
    } else if flags.cli_mode {
        println('Running in CLI mode...')
        exit(0)
    }

    // Profile.
    if os.exists(flags.config_path) {
        profile := toml.parse_text(os.read_file(flags.config_path) or {""})!
    } else {
        eprintln('Configuration file not found at path: ${flags.config_path}')
    }

    // Test Templating.
    mut dtmi := dtm.initialize()

    defer {
        dtmi.stop_cache_handler()
    }

    mut theme_test := map[string]dtm.DtmMultiTypeMap{}

    theme_test['timeout'] = 30
    theme_test['user']    = osutils.username()!
    theme_test['ver']     = osutils.grub_version()!
    theme_test['pkg']     = osutils.packages()!
    theme_test['os']      = osutils.operating_system()!
    theme_test['cpu']     = osutils.cpu_cores()!
    theme_test['mem']     = osutils.memory()!

    render_theme_test := dtmi.expand('theme.txt', placeholders: &theme_test)

    // Write to file.
    os.write_file('theme.txt', render_theme_test) or {
        eprintln('Error writing theme file: ${err}')
        exit(1)
    }
}