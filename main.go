package main

import (
	"fmt"
	"os"

	"github.com/lonepeon/tailog/internal/cmd"
)

func main() {
	if err := cmd.Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
