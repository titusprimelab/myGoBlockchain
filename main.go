package main

import (
	"os"
	"github.com/tkivite/myGoBlockchain/cli"
)
func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
