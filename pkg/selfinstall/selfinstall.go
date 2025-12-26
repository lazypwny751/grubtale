package selfinstall

import (
	"embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Install(serviceFiles embed.FS, configPath, grubPath, initSystem string) error {
	// 1. Detect GRUB path
	finalGrubPath := grubPath
	if finalGrubPath == "" {
		possiblePaths := []string{"/boot/grub", "/boot/grub2"}
		for _, p := range possiblePaths {
			if _, err := os.Stat(p); err == nil {
				finalGrubPath = p
				break
			}
		}
	}

	if finalGrubPath == "" {
		return fmt.Errorf("GRUB directory not found. Please specify with -grub-path")
	}
	slog.Info("Found GRUB directory", "path", finalGrubPath)

	// 2. Install binary to /usr/local/bin
	binPath := "/usr/local/bin"
	if err := os.MkdirAll(binPath, 0755); err != nil {
		return fmt.Errorf("failed to create %s: %w", binPath, err)
	}

	selfPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	destBin := filepath.Join(binPath, "grubtale")
	// Copy binary instead of symlink to ensure it works even if source is moved/deleted
	if err := copyFile(selfPath, destBin); err != nil {
		return fmt.Errorf("failed to install binary to %s: %w", destBin, err)
	}
	if err := os.Chmod(destBin, 0755); err != nil {
		return fmt.Errorf("failed to chmod binary: %w", err)
	}
	slog.Info("Installed binary", "path", destBin)

	// 3. Setup configuration
	configDir := "/etc/grubtale"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config dir %s: %w", configDir, err)
	}

	if configPath != "" {
		destConfig := filepath.Join(configDir, "grubtale.json")
		if err := copyFile(configPath, destConfig); err != nil {
			return fmt.Errorf("failed to copy config file: %w", err)
		}
		slog.Info("Installed configuration", "path", destConfig)
	}

	// 4. Install Service
	if initSystem == "auto" {
		if isSystemd() {
			initSystem = "systemd"
		} else if isSysVinit() {
			initSystem = "sysvinit"
		} else {
			return fmt.Errorf("could not detect init system, please specify with -init-system")
		}
	}

	switch initSystem {
	case "systemd":
		if err := installSystemd(serviceFiles); err != nil {
			return err
		}
	case "sysvinit":
		if err := installSysVinit(serviceFiles); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported init system: %s", initSystem)
	}

	slog.Info("Installation completed successfully!")
	return nil
}

func isSystemd() bool {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return true
	}
	// Check if pid 1 is systemd
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	out, err := cmd.Output()
	if err == nil && strings.TrimSpace(string(out)) == "systemd" {
		return true
	}
	return false
}

func isSysVinit() bool {
	if _, err := os.Stat("/etc/init.d"); err == nil {
		return true
	}
	return false
}

func installSystemd(fs embed.FS) error {
	content, err := fs.ReadFile("grubtale.service")
	if err != nil {
		return fmt.Errorf("failed to read embedded service file: %w", err)
	}

	servicePath := "/etc/systemd/system/grubtale.service"
	if err := os.WriteFile(servicePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	// Reload daemon and enable service
	exec.Command("systemctl", "daemon-reload").Run()
	if err := exec.Command("systemctl", "enable", "grubtale.service").Run(); err != nil {
		return fmt.Errorf("failed to enable systemd service: %w", err)
	}

	slog.Info("Installed systemd service", "path", servicePath)
	return nil
}

func installSysVinit(fs embed.FS) error {
	content, err := fs.ReadFile("grubtale.sh")
	if err != nil {
		return fmt.Errorf("failed to read embedded script: %w", err)
	}

	scriptPath := "/etc/init.d/grubtale"
	if err := os.WriteFile(scriptPath, content, 0755); err != nil {
		return fmt.Errorf("failed to write init script: %w", err)
	}

	// Try to enable service based on distro
	if _, err := exec.LookPath("update-rc.d"); err == nil {
		exec.Command("update-rc.d", "grubtale", "defaults").Run()
	} else if _, err := exec.LookPath("chkconfig"); err == nil {
		exec.Command("chkconfig", "--add", "grubtale").Run()
	}

	slog.Info("Installed SysVinit script", "path", scriptPath)
	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
