# Grubtale 🎮💀

An Undertale-inspired GRUB theme that fills your boot process with **DETERMINATION**.

![Grubtale Preview](https://github.com/user-attachments/assets/b626ba17-878d-412a-90c8-a6c5c12e0509)

## Features ✨

- **Dynamic Background Generation**: Each boot shows different Undertale-inspired quotes adapted for Linux
- **Real-time System Info**: Displays kernel version, package count, and uptime on boot screen
- **Pixel-Perfect Fonts**: Uses authentic PixelOperator fonts for that retro game feel
- **Multiple Background Options**: Choose between castle and home themes
- **Easy Installation**: One-command installation with automatic GRUB configuration
- **Service Integration**: Optional systemd/SysVInit services for automatic updates
- **Smart Background Generator**: Creates dynamic backgrounds with system information

## Screenshots 📸

*Despite everything, it's still you... booting Linux.*

## Installation 🚀

### Quick Installation
```bash
# Make the script executable
chmod +x grubtale.sh

# Preview the theme
./grubtale.sh preview

# Install the theme (requires root)
sudo ./grubtale.sh install
```

### Available Commands
```bash
./grubtale.sh help                    # Show help and random quote
./grubtale.sh preview                 # Generate and show theme preview
./grubtale.sh generate                # Generate background image only
./grubtale.sh status                  # Check service status
sudo ./grubtale.sh install            # Install the Grubtale theme and services
sudo ./grubtale.sh uninstall          # Remove the Grubtale theme and services
sudo ./grubtale.sh install-services   # Install only background update services
sudo ./grubtale.sh uninstall-services # Remove only background update services
```

## What the Script Does 🔧

The `grubtale.sh` script provides a complete theme management system:

### Background Generation
- Uses ImageMagick to overlay text on background images
- Displays random Undertale-inspired Linux quotes
### Background Generation
- Uses ImageMagick to overlay text on background images with proper font sizing
- Displays random Undertale-inspired Linux quotes
- Shows real-time system information (hostname, kernel, packages, uptime, date)
- Creates unique backgrounds for each preview/installation
- Uses theme.txt specifications for consistent styling and positioning
- Improved font sizes and positioning for better readability

### Installation Process
1. **Dependency Check**: Verifies ImageMagick is installed
2. **Background Generation**: Creates dynamic background with enhanced system info
3. **File Deployment**: Copies theme files to `/boot/grub/themes/grubtale/`
4. **GRUB Configuration**: Updates `/etc/default/grub` automatically
5. **GRUB Rebuild**: Runs `update-grub` to apply changes
6. **Service Installation**: Automatically installs systemd services for background updates

### Features
- **35+ Undertale-Linux Quotes**: Clever adaptations like "You're gonna have a bad boot"
- **Colorful ASCII Art Logo**: Beautiful GRUBTALE banner with random quotes
- **Safety Backups**: Backs up original GRUB config before changes
- **Smart Detection**: Auto-detects and uses available image viewers for preview
- **Error Handling**: Comprehensive error checking and user guidance
- **Automatic Updates**: Background refreshes on boot and daily with current system info
- **Service Management**: Install/uninstall/status commands for system services

## System Services 🔧

The Grubtale theme now includes automatic background update services that refresh the background with current system information on every boot and daily.

### Automatic Installation
When you run `sudo ./grubtale.sh install`, services are automatically installed and configured.

### Manual Service Management
```bash
# Install only services (without theme)
sudo ./grubtale.sh install-services

# Remove only services
sudo ./grubtale.sh uninstall-services

# Check service status
./grubtale.sh status
```

### Service Details
- **Boot Updates**: Background refreshes 2 minutes after each boot
- **Daily Updates**: Background refreshes daily with fresh system information
- **SystemD Integration**: Uses modern systemd timers and services
- **Logging**: All service activity logged to systemd journal

## Customization 🎨

### Change Background Image
Edit `grubtale.sh` and modify the `BACKGROUND_IMAGE` variable:
```bash
BACKGROUND_IMAGE="thome.jpg"  # or "castle.png"
```

### Add Your Own Quotes
Edit `grubtale.sh` and add to the `QUOTATION` array:
```bash
QUOTATION=(
    "Your custom quote here"
    # ... existing quotes
)
```

### Modify Theme Colors
Edit `theme.txt` to customize colors, positions, and styling.

## Requirements 📋

- **Linux Distribution** with GRUB2
- **ImageMagick** (`sudo apt install imagemagick`)
- **Root access** for installation
- **Font support** in GRUB (usually automatic)

## File Structure 📁

```
grubtale/
├── grubtale.sh           # Main installation script
├── theme.txt             # GRUB theme configuration
├── fonts/                # PixelOperator fonts
│   ├── PixelOperator.ttf
│   ├── PixelOperator8.ttf
│   └── PixelOperatorHB.ttf
├── images/               # Background images
│   ├── castle.png
│   ├── thome.jpg
│   └── ref.txt
├── services/             # System service files
│   ├── grubtale.service
│   ├── grubtale.timer
│   └── grubtale-sysvinit
├── assets/               # Desktop integration
│   ├── grubtale.desktop
│   └── icon.png
├── LICENSE               # MIT License
└── README.md             # This file
```

## Troubleshooting 🔧

### Theme Not Showing
1. Check if GRUB theme is enabled: `grep GRUB_THEME /etc/default/grub`
2. Verify theme files exist: `ls -la /boot/grub/themes/grubtale/`
3. Run `sudo update-grub` and reboot

### Fonts Not Loading
- Ensure fonts are in the correct directory
- Check font permissions: `sudo chmod -R 644 /boot/grub/themes/grubtale/fonts/`

### Background Not Generating
- Install ImageMagick: `sudo apt install imagemagick`
- Check image paths in the script
- Verify font files are accessible

## Contributing 🤝

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Credits 🙏

- **Undertale** by Toby Fox - Original game and inspiration
- **PixelOperator Font** - Retro pixel font family
- **Background Images** - Community artwork (see `images/ref.txt`)
- **Linux Community** - For keeping the spirit of determination alive

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact 📧

- **Discord**: lazypwny751
- **Twitter**: Ahmetta02120401s

---

*"Despite everything, it's still GNU/Linux."*

**Stay determined!** ⭐

