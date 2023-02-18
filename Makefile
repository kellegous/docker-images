all: bin/build

bin/buildimg:
	go build -o $@ github.com/kellegous/buildimg

bin/build: main.go bin/buildimg
	go build -o $@

clean:
	rm -rf bin