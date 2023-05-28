package processnames

import (
	"bufio"
	"errors"
	"fmt"
	"name-counter/pkg/readercounter"
	"os"
	"sort"
)

func ProcessNames(path string) (map[string]int, error) {
	var scanner *bufio.Scanner
	if path == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(path)
		defer func() {
			_ = file.Close()
		}()
		if err != nil {
			return nil, err
		}
		scanner = bufio.NewScanner(file)
	}
	return readercounter.ReadCountNames(scanner), nil
}

func PrintNames(names map[string]int) {
	for name, freq := range names {
		fmt.Printf("%s: %v\n", name, freq)
	}
}

func PrintNamesConditional(names map[string]int, cond string) error {
	length := len(names)
	namesArr := make([]struct {
		Name string
		Freq int
	}, length)
	i := 0
	for name, freq := range names {
		namesArr[i].Name = name
		namesArr[i].Freq = freq
		i++
	}
	sort.Slice(namesArr, func(i, j int) bool { return namesArr[i].Freq < namesArr[j].Freq })
	switch cond {
	case "asc":
		for index := 0; index < length; index++ {
			fmt.Printf("%v: %v\n", namesArr[index].Name, namesArr[index].Freq)
		}
	case "desc":
		for index := length - 1; index > -1; index-- {
			fmt.Printf("%v: %v\n", namesArr[index].Name, namesArr[index].Freq)
		}
	default:
		return errors.New("error: unknown comparison condition")
	}
	return nil
}
