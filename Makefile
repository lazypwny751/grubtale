SRC := src/main.go
OUT := grubtale

all: build

build:
	go build -o $(OUT) $(SRC)

clean:
	rm -f $(OUT)

.PHONY: all build clean