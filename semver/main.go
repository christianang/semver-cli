package main

import (
	"log"
	"os"

	"github.com/christianang/semver-cli/cmd"
)

func main() {
	command := os.Args[1]
	args := os.Args[2:]

	commandSet := cmd.Set{}
	commandSet["compare"] = cmd.NewCompare(os.Stdout, os.Stderr)
	commandSet["help"] = cmd.NewHelp(os.Stdout, os.Stderr)

	err := commandSet[command].Execute(args)
	if err != nil {
		log.Fatal(err)
	}
}
