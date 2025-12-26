#!/bin/sh

CONFIG="/etc/grubtale/grubtale.json"
BIN="/usr/local/bin/grubtale"
OUTPUT="/boot/grub/themes/Grubtale"

case "${1}" in
	"start")
		if [ -f "$CONFIG" ]; then
			"$BIN" -config "$CONFIG"
			if [ -d "$OUTPUT" ]; then
				cp "$CONFIG" "$OUTPUT/grubtale.json"
			fi
		else
			"$BIN"
		fi
	;;
	*)
		echo "Usage: ${0##*/} {start}"
		exit 1
	;;
esac
