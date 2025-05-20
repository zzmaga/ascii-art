package main

import (
	"ascii-art/operations"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	arg := strings.Split(os.Args[1], "\n")
	banner := "standard"
	text := operations.GetText(arg, banner)

	fmt.Print(text)
}
