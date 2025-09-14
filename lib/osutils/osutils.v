module osutils

import os

pub fn username() !string {
	return os.getenv('USER')
}

pub fn grub_version() !string {
	if os.execute('grub-install --version').exit_code == 0 {
		return os.execute('grub-install --version').output.trim_space().split(' ').last().split('.')[0]
	} else {
		return 'nil'
	}
}

pub fn cpu_cores() !string {
	if os.exists("/proc/cpuinfo") {
		cores := os.read_file("/proc/cpuinfo")!.split_into_lines().filter(it.starts_with("cpu cores"))[0].split(":")[1].trim_space().int()
		return '${cores} Core'
	} else {
		return 'nil'
	}
}

pub fn memory() !string {
	if os.exists("/proc/meminfo") {
		mem_gib := f64(os.read_file("/proc/meminfo")!.split_into_lines().filter(it.starts_with("MemTotal"))[0].split(":")[1].trim_space().split(" ")[0].int()) / (1024*1024)
		return '${mem_gib:.1f} GiB'
	} else {
		return 'nil'
	}
}

pub fn operating_system() !string {
	mut name := 'nil'
	
	if os.exists('/etc/os-release') {
		content := os.read_file('/etc/os-release') or { '' }
		for line in content.split_into_lines() {
			if line.starts_with('NAME=') {
				name = line.replace('NAME=', '').trim('"')

				match name.to_lower() {
					'linux mint' { name = 'Mint' }
					'arch linux' { name = 'Arch' }
					'ubuntu' { name = 'Ubuntu' }
					'fedora' { name = 'Fedora' }
					'opensuse leap' { name = 'openSUSE' }
					'suse linux enterprise server' { name = 'SUSE' }
					else {}
				}
				break
			}
		}
	}

	return name
}

fn package_manager() string {
	mut pkg := 'unknown'
	mut distro_ids := []string{}

	if os.exists('/etc/os-release') {
		content := os.read_file('/etc/os-release') or { '' }
		for line in content.split_into_lines() {
			if line.starts_with('ID=') {
				distro_ids << line.replace('ID=', '').trim('"')
			}
			if line.starts_with('ID_LIKE=') {
				distro_ids << line.replace('ID_LIKE=', '').trim('"').split(' ')
			}
		}
	}

	for id in distro_ids {
		match id.to_lower() {
			'debian', 'ubuntu', 'linuxmint' { pkg = 'apt'; break }
			'arch', 'manjaro' { pkg = 'pacman'; break }
			'fedora', 'rhel', 'centos' { pkg = 'dnf'; break }
			'opensuse', 'suse' { pkg = 'zypper'; break }
			else {}
		}
	}

	return pkg
}

pub fn packages() !string {
	mut count := "nil"

	match package_manager() {
		'apt' {
			count = os.execute("dpkg -l | wc -l").output.trim_space()
		}
		'pacman' {
			count = os.execute("pacman -Q | wc -l").output.trim_space()
		}
		'dnf' {
			count = os.execute("dnf repoquery --pkgnarrow=all | wc -l").output.trim_space()
		}
		'zypper' {
			count = os.execute("zypper se -i | wc -l").output.trim_space()
		} else {
			count = "nil"
		}
	}

	return count
}