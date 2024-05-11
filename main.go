package main

import (
	"flag"
)

var host = flag.Bool("server", false, "Host via ssh")

func main() {
	flag.Parse()

	if *host {
		sshServer()
		return
	}

	runUI()
}
