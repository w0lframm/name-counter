package main

import (
	"fmt"
	"name-counter/cli"
	"os"
)

func main() {
	err := clicounter.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
