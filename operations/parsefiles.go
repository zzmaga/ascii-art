package operations

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFiles(banner string) map[rune][]string {
	file, err := os.Open(banner + ".txt")
	if err != nil {
		fmt.Println("Error while opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parsedData := make(map[rune][]string)

	ascii := 32

	for scanner.Scan() {
		var lines []string
		for i := 0; i < 8; i++ {
			lines = append(lines, scanner.Text())
			if i < 7 {
				scanner.Scan()
			}
		}
		parsedData[rune(ascii)] = lines
		ascii++
		scanner.Scan()
	}
	return parsedData
}
