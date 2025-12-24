SRC := main.go
OUT := grubtale

all: $(OUT)

$(OUT):
	go build -o $(OUT) $(SRC)

clean:
	rm -f $(OUT)

test:
	go test ./...

.PHONY: all $(OUT) clean
