all: build

prepare:
	go mod tidy

build: prepare
	mkdir -p ./bin
	GO111MODULE=on go build -o ./bin/bbs .

docker:
	docker build --network host --tag moncheeta/bbs .

run: build
	./bin/bbs

clean:
	rm -rf ./tmp
	rm -rf ./bin

PORT=8000

serve: build
	# ./bin/bbs --server
	ttyd -p $(PORT) -W -t "titleFixed=Damian's BBS" ./bin/bbs

