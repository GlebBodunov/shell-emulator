package main

import (
	"os"

	"shellemu/src/shell"
)

func main() {
	sh := shell.New(os.Stdout)
	sh.RunInteractive(os.Stdin)
}
