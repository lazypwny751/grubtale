#ifndef GRUBTALE_H
#define GRUBTALE_H

struct ImageDimensions {
    int width;
    int height;
};

int resize_image(const char *infile, const char *outfile, struct ImageDimensions new_size);
void help_message(void);

#endif // GRUBTALE_H