SRC := main.go
OUT := grubtale

all: $(OUT)

$(OUT):
	go build -o $(OUT) $(SRC)

clean:
	rm -f $(OUT)

test:
	go test ./...

install: $(OUT)
	sudo ./$(OUT) -install

.PHONY: all $(OUT) clean test install
