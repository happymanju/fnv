package main

import (
	"os"

	"github.com/happymanju/fnv/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
