package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(Calc(s))
}

func Calc(s string) int {
	var n int
	for _, rune := range s {
		if rune == '1' {
			n++
		}
	}
	return n
}
