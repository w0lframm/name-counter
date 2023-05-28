package clicounter

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func Start() error {
	path := flag.String("in", "", "A path to the file with names")
	flag.Parse()
	if *path == "" || len(flag.Args()) != 0 {
		func() {
			_, _ = fmt.Fprintf(os.Stderr, "Usage %v --in <path-to-file>\n", os.Args[0])
			flag.PrintDefaults()
		}()
		return errors.New("error: invalid command line")
	}
	fmt.Printf("Got path: %v\n", *path)
	names, err := readCount(*path)
	printNames(names)
	//reader_counter.ReadAndCount("C:\\Tmp\\name_counter\\names\\names.txt")
	return err
}

func readCount(path string) (map[string]int, error) {
	var scanner *bufio.Scanner
	var resErr error
	if path == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(path)
		defer func() {
			_ = file.Close()
		}()
		resErr = err
		fmt.Printf("file %v opened successfully!!\n", file.Name())
		scanner = bufio.NewScanner(file)
	}
	names := make(map[string]int, 1000)
	for scanner.Scan() {
		names[scanner.Text()]++
	}
	return names, resErr
}

func printNames(names map[string]int) {
	for name, freq := range names {
		fmt.Printf("%s: %v\n", name, freq)
	}
	fmt.Println("done printing!")
}
