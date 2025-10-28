#include <cairo.h>
#include <stdio.h>
#include <stdlib.h>

#include "grubtale.h"

void help_message(void) {
	printf("Grubtale - undertale Grub theme.\n");
	printf("Usage: grubtale -m <multiplier> -o <output_directory>\n");
	printf("Options:\n");
	printf("  --multiplier, -m <multiplier>        Specify the multiplier number\n");
	printf("  --output, -o <output_directory>      Specify the output image directory\n");
	printf("  --help, -h                           Show this help message\n");
}