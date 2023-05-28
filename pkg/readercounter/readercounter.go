package readercounter

import (
	"bufio"
)

func ReadCountNames(scanner *bufio.Scanner) map[string]int {
	names := make(map[string]int, 1500)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			names[text]++
		}
	}
	return names
}
