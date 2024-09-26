package main

import (
	"expense-tracker/pkg/cli"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("expected a command (add, list, summary, delete)")
	}

	cli.HandleCommand(os.Args[1:])
}
