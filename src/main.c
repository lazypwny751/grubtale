#include <stdio.h>
#include <stdlib.h>
#include <getopt.h>

#include "grubtale.h"

int main(int argc, char **argv) {
	int opt;
	int multiplier = 1;
	char *output_directory = ".";

	static struct option long_options[] = {
		{"multiplier", required_argument, 0, 'm'},
		{"output", required_argument, 0, 'o'},
		{"help", no_argument, 0, 'h'},
		{0, 0, 0, 0}
	};

	while ((opt = getopt_long(argc, argv, "hm:o:", long_options, NULL)) != -1) {
		switch (opt) {
			case 'm': {
				multiplier = atoi(optarg);
				break;
			}
			case 'o': {
				output_directory = optarg;
				break;
			}
			case 'h': {
				help_message();
				return EXIT_SUCCESS;
			}
			default: {
				fprintf(stderr, "Unknown option \"%s\", please type \"%s --help\" for get help about this program.\n", argv[optind], argv[0]);
				return EXIT_FAILURE;
			}
		}
	}

	printf("Multiplier: %d\n", multiplier);
	printf("Output Directory: %s\n", output_directory);

	return 0;
}