#!/bin/bash

#################################
# Undertale concept grub theme. #
#################################
# idk how can i make a conceptual approach for this theme,
# if you have an idea or something you can find me at:
#	discord: lazypwny751
#	twitter: Ahmetta02120401s

set -e

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

check_dependencies() {
    local deps=("convert" "identify")
    local missing=()
    
    for dep in "${deps[@]}"; do
        if ! command -v "$dep" &> /dev/null; then
            missing+=("$dep")
        fi
    done
    
    if [[ ${#missing[@]} -gt 0 ]]; then
        echo -e "${RED}Missing dependencies: ${missing[*]}${NC}"
        echo -e "${YELLOW}Please install ImageMagick: sudo apt install imagemagick${NC}"
        exit 1
    fi
}

generate_background() {
    local input_image="$1"
    local output_image="$2"
    
    echo -e "${BLUE}Generating background image...${NC}"
    
    # Get system info
    local packages=$(dpkg-query -f '${binary:Package}\n' -W 2>/dev/null | wc -l || echo "Unknown")
    local kernel=$(uname -r)
    local uptime=$(uptime -p 2>/dev/null || echo "Unknown")
    local quote="${QUOTATION[$RANDOM % ${#QUOTATION[@]}]}"
    
    # Generate the background with overlay text
    convert "images/${input_image}" \
        -font "fonts/PixelOperator.ttf" \
        -pointsize 32 -fill white -stroke black -strokewidth 1 \
        -gravity North -annotate +0+60 "$quote" \
        -pointsize 20 -fill yellow \
        -gravity Southwest -annotate +20+120 "Kernel: $kernel" \
        -gravity Southwest -annotate +20+90 "Packages: $packages" \
        -gravity Southwest -annotate +20+60 "Uptime: $uptime" \
        -gravity Southeast -annotate +20+30 "* You feel determined." \
        "$output_image"
    
    echo -e "${GREEN}Background generated: $output_image${NC}"
}

create_theme_structure() {
    local theme_dir="$1"
    
    echo -e "${BLUE}Creating theme structure in $theme_dir${NC}"
    
    mkdir -p "$theme_dir"
    
    # Copy fonts
    cp -r fonts/ "$theme_dir/"
    
    # Copy and update theme.txt
    sed "s/background.png/background.png/" theme.txt > "$theme_dir/theme.txt"
    
    echo -e "${GREEN}Theme structure created${NC}"
}

install_theme() {
    echo -e "${BLUE}Installing Grubtale theme...${NC}"
    
    # Check if running as root
    if [[ $EUID -ne 0 ]]; then
        echo -e "${RED}This operation requires root privileges.${NC}"
        echo -e "${YELLOW}Please run: sudo $0 install${NC}"
        exit 1
    fi
    
    # Create directories
    mkdir -p "$INSTALL_DIR"
    
    # Generate background
    generate_background "$BACKGROUND_IMAGE" "background.png"
    
    # Create theme structure
    create_theme_structure "$INSTALL_DIR"
    
    # Copy background
    cp background.png "$INSTALL_DIR/"
    
    # Update GRUB configuration
    if [[ -f "/etc/default/grub" ]]; then
        echo -e "${BLUE}Updating GRUB configuration...${NC}"
        
        # Backup original config
        cp /etc/default/grub /etc/default/grub.backup.$(date +%Y%m%d_%H%M%S)
        
        # Update or add theme line
        if grep -q "GRUB_THEME=" /etc/default/grub; then
            sed -i "s|^GRUB_THEME=.*|GRUB_THEME=\"$INSTALL_DIR/theme.txt\"|" /etc/default/grub
        else
            echo "GRUB_THEME=\"$INSTALL_DIR/theme.txt\"" >> /etc/default/grub
        fi
        
        # Update GRUB
        update-grub
        
        echo -e "${GREEN}Grubtale theme installed successfully!${NC}"
        echo -e "${YELLOW}Reboot to see the theme in action.${NC}"
    else
        echo -e "${RED}GRUB configuration file not found!${NC}"
        exit 1
    fi
}

uninstall_theme() {
    echo -e "${BLUE}Uninstalling Grubtale theme...${NC}"
    
    if [[ $EUID -ne 0 ]]; then
        echo -e "${RED}This operation requires root privileges.${NC}"
        echo -e "${YELLOW}Please run: sudo $0 uninstall${NC}"
        exit 1
    fi
    
    # Remove theme directory
    if [[ -d "$INSTALL_DIR" ]]; then
        rm -rf "$INSTALL_DIR"
        echo -e "${GREEN}Theme files removed${NC}"
    fi
    
    # Restore GRUB config
    if [[ -f "/etc/default/grub" ]]; then
        sed -i '/^GRUB_THEME=/d' /etc/default/grub
        update-grub
        echo -e "${GREEN}GRUB configuration restored${NC}"
    fi
    
    echo -e "${GREEN}Grubtale theme uninstalled successfully!${NC}"
}

preview_theme() {
    echo -e "${BLUE}Generating preview...${NC}"
    generate_background "$BACKGROUND_IMAGE" "preview.png"
    echo -e "${GREEN}Preview generated: preview.png${NC}"
    
    if command -v feh &> /dev/null; then
        feh preview.png &
    elif command -v eog &> /dev/null; then
        eog preview.png &
    else
        echo -e "${YELLOW}Install 'feh' or 'eog' to auto-open the preview${NC}"
    fi
}

show_help() {
    print_logo
    echo -e "${CYAN}Usage: $0 [OPTION]${NC}"
    echo
    echo -e "${GREEN}Options:${NC}"
    echo -e "  ${YELLOW}install${NC}     Install the Grubtale theme (requires root)"
    echo -e "  ${YELLOW}uninstall${NC}   Remove the Grubtale theme (requires root)"
    echo -e "  ${YELLOW}preview${NC}     Generate a preview of the theme"
    echo -e "  ${YELLOW}generate${NC}    Generate background image only"
    echo -e "  ${YELLOW}help${NC}        Show this help message"
    echo
    echo -e "${GREEN}Examples:${NC}"
    echo -e "  $0 preview"
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
        "preview")
            OPTION="preview"
            shift
            ;;
        "generate")
            OPTION="generate"
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
        check_dependencies
        install_theme
        ;;
    "uninstall")
        print_logo
        uninstall_theme
        ;;
    "preview")
        print_logo
        check_dependencies
        preview_theme
        ;;
    "generate")
        print_logo
        check_dependencies
        generate_background "$BACKGROUND_IMAGE" "background.png"
        ;;
    "help"|*)
        show_help
        ;;
esac
