package main

import (
	"errors"
	"flag"
	"fmt"
	"name-counter/pkg/processnames"
	"os"
)

func main() {
	path := flag.String("in", "", "A path to the file with names")
	cond := flag.String("sort", "", "Sort condition: ascending (asc) or descending (desc)")
	flag.Parse()
	if *path == "" || len(flag.Args()) != 0 {
		checkError(errors.New("error: invalid command line"))
	}

	names, err := processnames.ProcessNames(*path)
	checkError(err)

	switch *cond {
	case "":
		processnames.PrintNames(names)
	default:
		checkError(processnames.PrintNamesConditional(names, *cond))
	}
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\nUsage %v --in <path-to-file>\n", err, os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}
