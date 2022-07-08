package main

// ! Esta clase ha sido creada para aprender loops i arrays
import "fmt"

func main() {
	var times int64
	fmt.Scan(&times)

	var sliceTet []int

	for i := 0; i < int(times); i++ {
		sliceTet = append(sliceTet, i)
	}

	fmt.Println(sliceTet)
	fmt.Println("lenght de la slice: ", len(sliceTet))

}
