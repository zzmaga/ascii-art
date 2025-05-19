package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		return
	}
	arg := strings.Split(os.Args[1], "\\n")
	banner := "standard"
	fmt.Println(arg, banner)
}

func GetText(arg []string, banner string) {
	doneData := ParseFiles(banner)

	for _, str := range arg {
		for _, char := range str {

		}
	}
}

func ParseFiles(banner string) map[int][]string {
	file, err := os.Open(banner + ".txt")
	if err != nil {
		fmt.Println("Error while opening file")
		os.Exit(0)
	}
	defer file.Close()
	parsedData := make(map[rune][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i := 0; i <= 855; i++ {

		}
	}
	return nil
}

func tester() {
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[:1]
	adder(&sli)
	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", arr, sli, &arr, len(arr), cap(arr))
	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", sli, sli, sli, len(sli), cap(sli))
	fmt.Printf("%#v, %T, [%p], [%d], [%d]\n", sli, sli, sli, len(sli), cap(sli))
}
func adder(sli *[]int) []int {
	*sli = append(*sli, 6)
	return *sli
}
