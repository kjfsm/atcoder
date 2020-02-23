package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	fmt.Println(Calc(a, b, c, x))
}

func Calc(a, b, c, x int) int {
	var n int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					n++
				}
			}
		}
	}

	return n
}
