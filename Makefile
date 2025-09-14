PREFIX := /boot/grub/themes
BINOUT := grubtale

all: build

run:
	./$(BINOUT) --cli

build:
	v . -o $(BINOUT)

clean:
	rm $(BINOUT)

.PHONY: all build
