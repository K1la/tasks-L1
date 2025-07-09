package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)
	runestr := []rune(str)
	for i := len(runestr) - 1; i >= 0; i-- {
		fmt.Print(string(runestr[i]))
	}
}
