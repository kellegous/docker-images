.PHONY: all clean

all: bin/build

bin/buildimg:
	GOBIN="$(CURDIR)/bin" go install github.com/kellegous/buildimg@latest

bin/build: cmd/build/main.go bin/buildimg
	go build -o $@ ./cmd/build

clean:
	rm -rf bin