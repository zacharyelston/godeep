# Makefile
.PHONY: build run test clean

build:
	go build -o godeep

run: build
	./godeep

test:
	go test ./...

clean:
	rm -f godeep

