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
- **Makefile Support**: Modern build system for easy management

## Screenshots 📸

*Despite everything, it's still you... booting Linux.*

## Installation 🚀

### Quick Installation with Makefile
```bash
# Clone the repository
git clone https://github.com/lazypwny751/grubtale.git
cd grubtale

# Check dependencies
make check-deps

# Preview the theme
make preview

# Install the theme (requires root)
sudo make install
```

### Alternative Installation Methods

**Using the script directly:**
```bash
./grubtale.sh preview
sudo ./grubtale.sh install
```

**Full installation with systemd service:**
```bash
sudo make full-install
```

## Usage 🎯

### Makefile Commands
```bash
make help              # Show available commands
make check-deps        # Check for required dependencies
make preview           # Generate and show theme preview
make generate          # Generate background image only
sudo make install      # Install the Grubtale theme
sudo make uninstall    # Remove the Grubtale theme
sudo make install-service    # Install systemd service
sudo make clean        # Clean generated files
make test              # Run basic tests
```

### Script Commands
```bash
# Show help and random quote
./grubtale.sh help

# Generate preview image
./grubtale.sh preview

# Generate background only
./grubtale.sh generate

# Install theme
sudo ./grubtale.sh install

# Uninstall theme
sudo ./grubtale.sh uninstall
```

## System Services 🔧

### SystemD (Modern Systems)
```bash
# Copy service files
sudo cp services/grubtale.service /etc/systemd/system/
sudo cp services/grubtale.timer /etc/systemd/system/

# Enable daily background updates
sudo systemctl enable grubtale.timer
sudo systemctl start grubtale.timer
```

### SysVInit (Legacy Systems)
```bash
# Copy init script
sudo cp services/grubtale-sysvinit /etc/init.d/grubtale
sudo chmod +x /etc/init.d/grubtale

# Enable service
sudo update-rc.d grubtale defaults
```

## Customization 🎨

### Change Background Image
Edit `grubtale.sh` and modify the `BACKGROUND_IMAGE` variable:
```bash
BACKGROUND_IMAGE="thome.jpg"  # or "castle.png"
```

### Add Your Own Quotes
Edit `scripts/quotes.sh` and add to the `QUOTATION` array:
```bash
export QUOTATION=(
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
├── grubtale.sh           # Main script
├── install.sh            # Quick installer
├── theme.txt             # GRUB theme configuration
├── fonts/                # PixelOperator fonts
├── images/               # Background images
├── scripts/              # Quote scripts
├── services/             # System service files
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

