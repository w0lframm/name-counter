package main

import (
	//"name-counter/cli"
	"name-counter/reader-counter"
)

func main() {
	//cli.Start()
	reader_counter.ReadAndCount("C:\\Tmp\\name_counter\\names\\names.txt")
}
