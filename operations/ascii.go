package operations

import (
	"strings"
)

func GetText(lines []string, banner string) string {
	parsedData := ParseFiles(banner)
	var result strings.Builder

	for lineIdx, line := range lines {
		outputLines := make([]string, 8)
		for _, char := range line {
			if art, ok := parsedData[char]; ok {
				for i := 0; i < 8; i++ {
					outputLines[i] += art[i]
				}
			}
		}
		for i := 0; i < 8; i++ {
			result.WriteString(outputLines[i] + "\n")
		}
		if lineIdx != len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

// func tester() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	sli := arr[:1]
// 	adder(&sli)
// 	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", arr, sli, &arr, len(arr), cap(arr))
// 	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", sli, sli, sli, len(sli), cap(sli))
// 	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", sli, sli, sli, len(sli), cap(sli))
// }
// func adder(sli *[]int) []int {
// 	*sli = append(*sli, 6)
// 	return *sli
// }
