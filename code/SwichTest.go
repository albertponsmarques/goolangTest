package main

import "fmt"

func main() {
	ans := 10

	// ! cases need to be the same type of variable
	switch ans {
	case 1, -1, 3:
		fmt.Println("this is one or -1 or 3, yeah pretty cool")
	case 2:
		fmt.Println("this is two")
	default:
		fmt.Println("default thing")
	}

	// you can make a swich without the variable in it
	// case var = x
}
