package generator

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}

func GetOSName() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Linux"
	}
	defer file.Close()

	var id, name, prettyName string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
		} else if strings.HasPrefix(line, "NAME=") {
			name = strings.Trim(strings.TrimPrefix(line, "NAME="), "\"")
		} else if strings.HasPrefix(line, "ID=") {
			id = strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
		}
	}

	if prettyName != "" {
		// Shorten if too long
		if len(prettyName) > 20 {
			return prettyName[:17] + "..."
		}
		return prettyName
	}
	if name != "" {
		return name
	}
	if id != "" {
		// Capitalize first letter
		if len(id) > 0 {
			id = strings.ToUpper(id[:1]) + id[1:]
		}
		return id
	}

	return "Linux"
}

func GetCPUInfo() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "Unknown CPU"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				cpu := strings.TrimSpace(parts[1])
				// Shorten the CPU name
				replacements := []string{
					"Intel(R)", "",
					"Core(TM)", "",
					"AMD", "",
					"Ryzen", "Ryzen",
					"Processor", "",
					"CPU", "",
					"12th Gen", "",
					"13th Gen", "",
					"14th Gen", "",
				}
				r := strings.NewReplacer(replacements...)
				cpu = r.Replace(cpu)

				// Remove frequency (e.g. @ 3.60GHz)
				if idx := strings.Index(cpu, "@"); idx != -1 {
					cpu = cpu[:idx]
				}

				return strings.TrimSpace(cpu)
			}
		}
	}
	return "Unknown CPU"
}

func GetPackageCount() int {
	// Try apt (Debian/Ubuntu)
	if _, err := exec.LookPath("dpkg"); err == nil {
		cmd := exec.Command("dpkg-query", "-f", "${binary:Package}\n", "-W")
		output, err := cmd.Output()
		if err == nil {
			return bytes.Count(output, []byte("\n"))
		}
	}

	// Try pacman (Arch)
	if _, err := exec.LookPath("pacman"); err == nil {
		cmd := exec.Command("pacman", "-Q")
		output, err := cmd.Output()
		if err == nil {
			return bytes.Count(output, []byte("\n"))
		}
	}

	// Try rpm (Fedora/RHEL)
	if _, err := exec.LookPath("rpm"); err == nil {
		cmd := exec.Command("rpm", "-qa")
		output, err := cmd.Output()
		if err == nil {
			return bytes.Count(output, []byte("\n"))
		}
	}

	// Try apk (Alpine)
	if _, err := exec.LookPath("apk"); err == nil {
		cmd := exec.Command("apk", "info")
		output, err := cmd.Output()
		if err == nil {
			return bytes.Count(output, []byte("\n"))
		}
	}

	return 0
}

func GetGrubTimeout() int {
	file, err := os.Open("/etc/default/grub")
	if err != nil {
		return -1
	}
	defer file.Close()

	lastVal := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "GRUB_TIMEOUT=") {
			valStr := strings.TrimPrefix(line, "GRUB_TIMEOUT=")
			valStr = strings.Trim(valStr, "\"")
			valStr = strings.Trim(valStr, "'")
			val, err := strconv.Atoi(valStr)
			if err == nil {
				lastVal = val
			}
		}
	}
	return lastVal
}
