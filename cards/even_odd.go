package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	findEvenOrOdd(intSlice)

}

func findEvenOrOdd(dataList []int) {

	for _, data := range dataList {
		if data%2 == 0 {
			fmt.Println("Even", data)
		} else {
			fmt.Println("Odd", data)

		}
	}

}
