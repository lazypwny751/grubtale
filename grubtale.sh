#!/bin/sh

case "${1}" in
	"start")
		:
	;;
	*)
		echo "Usage: ${0##*/} {start}"
		exit 1
	;;
esac
