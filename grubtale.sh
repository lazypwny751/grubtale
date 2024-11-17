#!/bin/bash

#################################
# Undertale concept grub theme. #
#################################
# idk how can i make a conceptual approach for this theme,
# if you have an idea or something you can find me at:
#	discord: lazypwny751
#	twitter: Ahmetta02120401s

set -e

export OPTION="help"

while (( "${#}" > 0 )) ; do
	*)
		shift
	;;
done

case "${OPTION,,}" in
	"help"|*)
		echo "This is help text."
	;;
esac

# Image generation:
# 	convert castle.png -font PixelOperator.ttf -pointsize 48 -fill white \
  # -gravity North -annotate +0+60 "${QUOTATION[$RANDOM % ${#QUOTATION[@]}]}" \
  # -gravity South -annotate +0+10 "Packages: $(dpkg-query -f '${binary:Package}\n' -W | wc -l)" output_image.png
