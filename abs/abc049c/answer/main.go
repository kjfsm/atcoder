package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if Calc(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func Calc(s string) bool {
	return loop(s, "S")
}

func loop(str string, state string) bool {
	switch state {

	case "S":
		switch {
		case str == "":
			return true
		case len(str) >= 3 && str[:3] == "dre":
			if loop(str[3:], "dre") {
				return true
			}
			fallthrough
		case len(str) >= 4 && str[:4] == "eras":
			if loop(str[4:], "eras") {
				return true
			}
		}

	case "dre":
		switch {
		case len(str) >= 2 && str[:2] == "am":
			if loop(str[2:], "am") {
				return true
			}
		}

	case "am":
		switch {
		case str == "":
			return true
		case len(str) >= 2 && str[:2] == "er":
			if loop(str[2:], "er") {
				return true
			}
			fallthrough
		default:
			if loop(str, "S") {
				return true
			}
		}

	case "er":
		if loop(str, "S") {
			return true
		}
	case "eras":
		switch {
		case len(str) >= 1 && str[:1] == "e":
			if loop(str[1:], "e") {
				return true
			}
			fallthrough
		case len(str) >= 2 && str[:2] == "er":
			if loop(str[2:], "er") {
				return true
			}
		}
	case "e":
		if loop(str, "S") {
			return true
		}
	}
	return false
}
