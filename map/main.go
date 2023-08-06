package main

import "fmt"

func main() {
	fmt.Println("=================== Maps In Go ======================")

	// var newColor map[string]int
	// newColor1 := make(map[string]int)

	// create new item
	// newColor1["red"] = 1

	//delete
	// delete(newColor1, "red")

	colors := map[string]int{"Red": 1, "Green": 2, "Blue": 3}

	printMapKeyValue(colors)

	// fmt.Println("Map Color :", colors)
}

func printMapKeyValue(m map[string]int) {
	for color, value := range m {
		fmt.Println("Color ", color, ":", value)
	}
}
