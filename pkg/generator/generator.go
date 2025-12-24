package generator

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			id := strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
			// Capitalize first letter
			if len(id) > 0 {
				id = strings.ToUpper(id[:1]) + id[1:]
			}
			if strings.Contains(strings.ToLower(id), "mint") {
				return "Mint"
			}
			return id
		}
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

	return 0
}
