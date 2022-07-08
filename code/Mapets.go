package main

import "fmt"

func main() {
	var mapet map[string]int = map[string]int{
		"apples":  5,
		"bananas": 8,
		"pears":   2,
	}

	mapet["apples"] = 8
	//adding
	mapet["mandarines"] = 9

	//deleting
	delete(mapet, "pears")

	//check if key exists
	val, ok := mapet["mandarines"]
	fmt.Println(val, ok)

	//also you can print the lenght of a map len()
	fmt.Println(mapet)
}
