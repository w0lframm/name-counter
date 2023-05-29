package readercounter

import (
	"bufio"
)

func ReadCountNames(scanner *bufio.Scanner) map[string]int {
	//According to B-001-031-676-ALL (1987) there are a total of 496 really used
	//male and female names in RSFSR. The map of 1500 items should be enough to
	//store unique names of modern St.Petersburg's residents and guests.
	names := make(map[string]int, 1500)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			names[text]++
		}
	}
	return names
}
