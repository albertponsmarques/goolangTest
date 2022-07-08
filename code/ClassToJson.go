package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Computer struct {
	Motherboard  string
	Processor    string
	GraphicsCard string
	RamGB        int
	Fans         []string
}

func main() {
	comp := Computer{
		Motherboard:  "gigabyte",
		Processor:    "intel i9",
		GraphicsCard: "gigabyte 3090",
		RamGB:        32,
		Fans:         []string{"back", "front", "top"},
	}

	comp2 := Computer{
		Motherboard:  "msi",
		Processor:    "ryzen 9",
		GraphicsCard: "msi rx 6800",
		RamGB:        32,
		Fans:         []string{"back", "front", "top", "liquid cooled"},
	}

	var list []Computer
	list = append(list, comp)
	list = append(list, comp2)

	file, _ := json.MarshalIndent(list, "", " ")
	_ = ioutil.WriteFile("pcs.json", file, 0644)

	if _, err := os.Stat("pcs.json"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No se ha podido crear el fichero")
	} else {
		fmt.Println("Fichero creado corrctamente")
	}
}
