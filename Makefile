all: build

prepare:
	go mod tidy

build: prepare
	mkdir -p ./bin
	GO111MODULE=on go build -o ./bin/bbs .

run: build
	./bin/bbs

serve: build
	# ./bin/bbs --server
	ttyd -W -t fontFamily='JetBrains' -t fontSize=22 -t 'theme={ "cursor": "#ffffff", "foreground": "#ffffff", "background": "#000000", "black": "#000000", "red": "#ff8059", "green": "#44bc44", "yellow": "#d0bc00", "blue": "#2fafff", "magenta": "#feacd0", "cyan": "#00d3d0", "white": "#bfbfbf", "brightBlack": "#595959", "brightRed": "#ef8b50", "brightGreen": "#70b900", "brightYellow": "#c0c530", "brightBlue": "#79a8ff", "brightMagenta": "#b6a0ff", "brightCyan": "#6ae4b9", "brightWhite": "#ffffff" }' ./bin/bbs

clean:
	rm -rf ./tmp
	rm -rf ./bin

