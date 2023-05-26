package reader_counter

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndCount(path string) {
	//
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v opened successfully!!\n", file.Name())
	defer file.Close()

	names := make(map[string]int, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names[scanner.Text()]++
	}
	printNames(names)
}

func printNames(names map[string]int) {
	for name, freq := range names {
		fmt.Printf("%s: %v\n", name, freq)
	}
	fmt.Println("done printing!")
}
