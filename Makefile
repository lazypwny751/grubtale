PREFIX := /boot/grub/themes
BINOUT := grubtale

all: build

build:
	$(CC) "src/main.c" "src/grubtale.c" -o "$(BINOUT)" `pkg-config --cflags --libs cairo`

clean:
	rm "$(BINOUT)"

.PHONY: all clean