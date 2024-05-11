all: build

prepare:
	go mod tidy

build: prepare
	mkdir -p ./bin
	GO111MODULE=on go build -o ./bin/bbs .

run: build
	./bin/bbs

PORT=8000

serve: build
	# ./bin/bbs --server
	ttyd -p $(PORT) -W ./bin/bbs

clean:
	rm -rf ./tmp
	rm -rf ./bin

