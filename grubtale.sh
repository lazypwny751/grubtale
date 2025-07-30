#!/bin/bash

#################################
# Undertale concept grub theme. #
#################################
# idk how can i make a conceptual approach for this theme,
# if you have an idea or something you can find me at:
#	discord: lazypwny751
#	twitter: Ahmetta02120401s

set -e

# Configuration
SERVICE_TYPE="systemd"  # Default service type (systemd or sysvinit)

# Undertale quotes adapted for Linux/GRUB
QUOTATION=(
    "You're gonna have a bad boot."
    "It's a beautiful day to compile. CPUs are running, fans are spinning…"
    "Do you wanna have a kernel panic?"
    "Despite everything, it's still GNU."
    "The shadow of GRUB looms above, filling you with determination."
    "But Linux refused."
    "Don't sudo rm -rf, and don't alias it, alright?"
    "You noob! In this world, it's root or be rooted."
    "Knock knock." "Who's there?" "Ls." "Ls who?" "Ls there any files in this directory?"
    "Nyeh heh heh! I've aliased rm to echo."
    "* The system fills you with determination."
    "* You check your processes. HP fully restored."
    "* You're blue now. That's my attack!"
    "* But nobody came... to /dev/null."
    "* In this world, it's kill -9 or be killed."
    "* You feel your configs crawling on your back."
    "* The power of package managers shines within you."
    "* You're filled with... BASH!"
    "* Would you like to install vim? (Y/n)"
    "* It's dangerous to code alone! Take this: man pages."
    "* sudo make me a sandwich"
    "* /home is where the heart is."
    "* The terminal whispers the names of forgotten commands."
    "* You gained 0 EXP and 0 gold. Your code compiled successfully!"
    "* Somewhere, a penguin is dancing."
    "GAME OVER. Insert coin or press any key to continue."
    "* The CPU fan spins quietly. You are filled with determination."
    "* Choose your shell wisely, young padawan."
    "* With great root access comes great responsibility."
    "* Error 404: Motivation not found. Please reboot yourself."
    "* Segfault happens. You feel bad time approaching."
    "* The kernel modules are loaded. You are ready for battle."
    "* Memory leaks everywhere. But you stay determined."
    "* In this world, it's compile or be deprecated."
    "* The filesystem whispers: 'It's a beautiful day to mount.'"
)

# Default values
OPTION="help"
GRUB_THEMES_DIR="/boot/grub/themes"
THEME_NAME="grubtale"
BACKGROUND_IMAGE="castle.png"
INSTALL_DIR="${GRUB_THEMES_DIR}/${THEME_NAME}"
SERVICE_TYPE="systemd"  # Default service type

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Functions
print_logo() {
    echo -e "${PURPLE}"
    echo "   ██████╗ ██████╗ ██╗   ██╗██████╗ ████████╗ █████╗ ██╗     ███████╗"
    echo "  ██╔════╝ ██╔══██╗██║   ██║██╔══██╗╚══██╔══╝██╔══██╗██║     ██╔════╝"
    echo "  ██║  ███╗██████╔╝██║   ██║██████╔╝   ██║   ███████║██║     █████╗  "
    echo "  ██║   ██║██╔══██╗██║   ██║██╔══██╗   ██║   ██╔══██║██║     ██╔══╝  "
    echo "  ╚██████╔╝██║  ██║╚██████╔╝██████╔╝   ██║   ██║  ██║███████╗███████╗"
    echo "   ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝╚══════╝╚══════╝"
    echo -e "${NC}"
    echo -e "${CYAN}An Undertale-inspired GRUB theme manager${NC}"
    echo -e "${YELLOW}\"${QUOTATION[$RANDOM % ${#QUOTATION[@]}]}\"${NC}"
    echo
}

generate_dynamic_theme() {
    local theme_dir="$1"
    
    echo -e "${BLUE}Generating dynamic theme configuration...${NC}"
    
    # Get minimal system info
    local packages=$(dpkg-query -f '${binary:Package}\n' -W 2>/dev/null | wc -l || echo "0")
    local os_info=$(lsb_release -d 2>/dev/null | cut -f2 || echo "GNU/Linux")
    local cpu_cores=$(nproc 2>/dev/null || echo "Unknown")
    local quote="${QUOTATION[$RANDOM % ${#QUOTATION[@]}]}"
    local current_date=$(date "+%B %d, %Y")
    
    # Truncate long strings for 720p display
    local os_short="${os_info:0:25}"
    [[ ${#os_info} -gt 25 ]] && os_short="${os_short}..."
    
    local quote_short="${quote:0:45}"
    [[ ${#quote} -gt 45 ]] && quote_short="${quote_short}..."
    
    # Generate clean and compact theme.txt for 720p (1280x720)
    cat > "$theme_dir/theme.txt" << EOF
# GRUB2 Theme - Grubtale (Compact) - 720p Format
# An Undertale-inspired GRUB theme with clean layout
# Generated on: $current_date

# Background settings
desktop-image: "background.png"
desktop-color: "#000000"

# Title - Top left
+ label {
  left = 30
  top = 30
  align = "left"
  font = "determination_mono 24"
  color = "#FFD700"
  text = "GRUBTALE"
}

# System info block - Bottom left
+ label {
  left = 30
  top = 620
  align = "left"
  font = "determination_mono 14"
  color = "#FFFFFF"
  text = "OS: $os_short"
}

+ label {
  left = 30
  top = 640
  align = "left"
  font = "determination_mono 14"
  color = "#FFFFFF"
  text = "Packages: $packages"
}

+ label {
  left = 30
  top = 660
  align = "left"
  font = "determination_mono 14"
  color = "#FFFFFF"
  text = "CPU Cores: $cpu_cores"
}

# Random quote - Top center
+ label {
  left = 640
  top = 30
  align = "center"
  font = "determination_mono 16"
  color = "#FFFF00"
  text = "$quote_short"
}

# Determination message - Bottom right
+ label {
  left = 1250
  top = 680
  align = "right"
  font = "determination_mono 14"
  color = "#FFD700"
  text = "* Stay determined."
}

# Update timestamp - Top right
+ label {
  left = 1250
  top = 10
  align = "right"
  font = "determination_mono 12"
  color = "#888888"
  text = "$current_date"
}

# Boot menu - Centered
+ boot_menu {
  left = 320
  top = 180
  width = 640
  height = 360
  item_font = "determination_mono 16"
  selected_item_font = "determination_mono 16"
  item_color = "#CCCCCC"
  selected_item_color = "#FFFFFF"
  item_height = 32
  item_padding = 8
  item_spacing = 4
  icon_width = 24
  icon_height = 24
  scrollbar = true
  scrollbar_width = 16
  scrollbar_thumb = "#FFD700"
  scrollbar_frame = "#444444"
}

# Numeric countdown - Bottom center (no progress bar)
+ label {
  id = "__timeout__"
  left = 640
  top = 580
  align = "center"
  font = "determination_mono 18"
  color = "#FFD700"
  text = "Auto boot in: %d"
}

# Font definitions
title-font: "determination_mono 18"
message-font: "determination_mono 16"
terminal-font: "determination_mono 14"
EOF
    
    echo -e "${GREEN}Dynamic theme configuration generated: $theme_dir/theme.txt${NC}"
}

create_theme_structure() {
    local theme_dir="$1"
    
    echo -e "${BLUE}Creating theme structure in $theme_dir${NC}"
    
    mkdir -p "$theme_dir"
    
    # Copy fonts
    cp -r fonts/ "$theme_dir/"
    
    # Copy background image directly
    cp "images/$BACKGROUND_IMAGE" "$theme_dir/background.png"
    
    # Generate dynamic theme.txt with current system info
    generate_dynamic_theme "$theme_dir"
    
    echo -e "${GREEN}Theme structure created${NC}"
}

install_service() {
    local service_type="$1"
    
    echo -e "${BLUE}Installing $service_type service...${NC}"
    
    # Create grubtale directory in /usr/local/share
    mkdir -p /usr/local/share/grubtale
    
    # Copy necessary files
    cp grubtale.sh /usr/local/share/grubtale/
    cp -r fonts/ /usr/local/share/grubtale/
    cp -r images/ /usr/local/share/grubtale/
    chmod +x /usr/local/share/grubtale/grubtale.sh
    
    # Create update script
    cat > /usr/local/share/grubtale/update-theme.sh << 'EOF'
#!/bin/bash
# Grubtale theme updater
cd /usr/local/share/grubtale
./grubtale.sh update-theme
EOF
    chmod +x /usr/local/share/grubtale/update-theme.sh
    
    if [[ "$service_type" == "systemd" ]]; then
        # Install systemd service
        cat > /etc/systemd/system/grubtale.service << EOF
[Unit]
Description=Grubtale Theme Updater
Documentation=https://github.com/lazypwny751/grubtale
After=network.target

[Service]
Type=oneshot
User=root
WorkingDirectory=/usr/local/share/grubtale
ExecStart=/usr/local/share/grubtale/update-theme.sh
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
EOF
        
        systemctl daemon-reload
        systemctl enable grubtale.service
        echo -e "${GREEN}SystemD service installed and enabled${NC}"
        
    elif [[ "$service_type" == "sysvinit" ]]; then
        # Install SysVInit script
        cp services/grubtale-sysvinit /etc/init.d/grubtale
        chmod +x /etc/init.d/grubtale
        
        # Enable service based on distribution
        if command -v update-rc.d &>/dev/null; then
            update-rc.d grubtale defaults
        elif command -v chkconfig &>/dev/null; then
            chkconfig --add grubtale
            chkconfig grubtale on
        fi
        
        echo -e "${GREEN}SysVInit service installed and enabled${NC}"
    fi
    
    echo -e "${YELLOW}Service will update theme on every boot${NC}"
}

uninstall_service() {
    echo -e "${BLUE}Uninstalling services...${NC}"
    
    # Stop and disable systemd service
    if systemctl is-enabled grubtale.service &>/dev/null; then
        systemctl disable grubtale.service 2>/dev/null || true
        systemctl stop grubtale.service 2>/dev/null || true
        rm -f /etc/systemd/system/grubtale.service
        systemctl daemon-reload
        echo -e "${GREEN}SystemD service removed${NC}"
    fi
    
    # Remove SysVInit service
    if [[ -f "/etc/init.d/grubtale" ]]; then
        if command -v update-rc.d &>/dev/null; then
            update-rc.d grubtale remove
        elif command -v chkconfig &>/dev/null; then
            chkconfig grubtale off
            chkconfig --del grubtale
        fi
        rm -f /etc/init.d/grubtale
        echo -e "${GREEN}SysVInit service removed${NC}"
    fi
    
    # Remove application directory
    rm -rf /usr/local/share/grubtale
    
    echo -e "${GREEN}All services uninstalled${NC}"
}

update_theme() {
    # This function is called by the service to update the theme
    if [[ -d "$INSTALL_DIR" ]]; then
        echo "Updating Grubtale theme..."
        generate_dynamic_theme "$INSTALL_DIR"
        update-grub 2>/dev/null || grub-mkconfig -o /boot/grub/grub.cfg 2>/dev/null || true
        echo "Grubtale theme updated successfully"
    fi
}

install() {
    check_root
    
    echo -e "${YELLOW}Installing Grubtale GRUB theme...${NC}"
    
    # Install theme
    create_theme_structure "$INSTALL_DIR"
    
    # Install service (default: systemd)
    install_service "$SERVICE_TYPE"
    
    # Update GRUB configuration to use the theme
    if grep -q "GRUB_THEME=" /etc/default/grub; then
        sed -i "s|^GRUB_THEME=.*|GRUB_THEME=\"$INSTALL_DIR/theme.txt\"|" /etc/default/grub
    else
        echo "GRUB_THEME=\"$INSTALL_DIR/theme.txt\"" >> /etc/default/grub
    fi
    
    # Update GRUB
    echo -e "${BLUE}Updating GRUB configuration...${NC}"
    update-grub 2>/dev/null || grub-mkconfig -o /boot/grub/grub.cfg
    
    echo -e "${GREEN}Grubtale theme installed successfully!${NC}"
    echo -e "${YELLOW}The theme will be automatically updated on every boot.${NC}"
    echo -e "${YELLOW}Reboot to see the new theme.${NC}"
}

uninstall() {
    check_root
    
    echo -e "${YELLOW}Uninstalling Grubtale GRUB theme...${NC}"
    
    # Uninstall services
    uninstall_service
    
    # Remove theme directory
    if [[ -d "$INSTALL_DIR" ]]; then
        rm -rf "$INSTALL_DIR"
        echo -e "${GREEN}Theme directory removed${NC}"
    fi
    
    # Remove GRUB_THEME from configuration
    if grep -q "GRUB_THEME=" /etc/default/grub; then
        sed -i '/^GRUB_THEME=/d' /etc/default/grub
        echo -e "${GREEN}GRUB configuration updated${NC}"
    fi
    
    # Update GRUB
    echo -e "${BLUE}Updating GRUB configuration...${NC}"
    update-grub 2>/dev/null || grub-mkconfig -o /boot/grub/grub.cfg
    
    echo -e "${GREEN}Grubtale theme uninstalled successfully!${NC}"
    echo -e "${YELLOW}Reboot to see the default GRUB theme.${NC}"
}

check_root() {
    if [[ $EUID -ne 0 ]]; then
        echo -e "${RED}This operation requires root privileges.${NC}"
        echo -e "${YELLOW}Please run: sudo $0 $1${NC}"
        exit 1
    fi
}

show_help() {
    print_logo
    echo -e "${CYAN}Usage: $0 [OPTION]${NC}"
    echo
    echo -e "${GREEN}Options:${NC}"
    echo -e "  ${YELLOW}install${NC}        Install the Grubtale theme with boot-time service (requires root)"
    echo -e "  ${YELLOW}uninstall${NC}      Remove the Grubtale theme and service (requires root)"
    echo -e "  ${YELLOW}update-theme${NC}   Update theme with current system info (for service use)"
    echo -e "  ${YELLOW}help${NC}           Show this help message"
    echo
    echo -e "${GREEN}Examples:${NC}"
    echo -e "  sudo $0 install"
    echo -e "  sudo $0 uninstall"
    echo
    echo -e "${PURPLE}\"Despite everything, it's still GNU/Linux.\"${NC}"
}

# Parse command line arguments
while (( "${#}" > 0 )) ; do
    case "${1,,}" in
        "install")
            OPTION="install"
            shift
            ;;
        "uninstall")
            OPTION="uninstall"
            shift
            ;;
        "update-theme")
            OPTION="update-theme"
            shift
            ;;
        "help"|"-h"|"--help")
            OPTION="help"
            shift
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            OPTION="help"
            shift
            ;;
    esac
done

# Main execution
case "${OPTION,,}" in
    "install")
        print_logo
        install
        ;;
    "uninstall")
        print_logo
        uninstall
        ;;
    "update-theme")
        update_theme
        ;;
    "help"|*)
        show_help
        ;;
esac
