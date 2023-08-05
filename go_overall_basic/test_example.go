package main

import (
	"fmt"
	"time"
)

func main() {

	/// variales

	fmt.Println("======== Basic Variables =============")

	a := 1
	b := "Messi"
	var c float32 = 1.2

	fmt.Printf("a is %d and b is %s and c is %f \n", a, b, c)

	//difference between array and slice is that array is fixed length and slice is dynamic length

	testSlice := []string{"hello", "hi", "namaste", "sawero"}
	// testArray := [3]int{1, 2, 3}
	// testMap := map[string]int{"a": 1, "b": 2}

	// for loop

	fmt.Println("======== Basic Loop =============")

	i := 1

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("======== ===== =============")

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("======== Basic Loop in Slice =============")

	for k := 0; k < len(testSlice); k++ {

		fmt.Println("Value of Slice :" + testSlice[k])
	}

	fmt.Println("======== Basic Loop with condition =============")

	for m := 0; m <= 5; m++ {
		if m%2 == 0 {
			// fmt.Printf("Continue hit %d \n", m)
			continue
		}
		fmt.Println(m)
	}

	fmt.Println("======== Loop With Range =============")

	for i, value := range testSlice {
		fmt.Printf("The index is %d and value is %s\n", i, value)
	}

	fmt.Println("======== Loop With Range For Map Data =============")

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// conditional if/else
	fmt.Println("======== Basic If/else conditions =============")

	if x := 7; x == 6 {
		fmt.Printf("The value of x is 6 \n")
	} else if x == 7 {
		fmt.Printf("The value of x is 7 \n")
	} else {
		fmt.Printf("The value of x is different, i.e %d\n", x)
	}

	// switch statement
	fmt.Println("======== Basic Switch statements =============")

	s := 2

	switch s {
	case 1:
		fmt.Println("The value is 1")

	case 2:
		fmt.Println("The value is 2")

	case 3:
		fmt.Println("The value is 3 ")

	default:
		fmt.Println("The condition is default")
	}

	fmt.Println("======== =========== =============")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// functions

	fmt.Println("======== Functions with single return type =============")

	result := addData(1, 2)
	fmt.Printf("The addition of 1+2 is %d \n", result)

	fmt.Println("======== Functions with multiple return type =============")
	// _, y := returnMulipelVal(10, 12)
	x, y := returnMulipelVal(10, 12)

	fmt.Println("The value of x and y is ", x, y)

	fmt.Println("======== Variadic functions  =============") // any number of trailing args

	total := sums(2, 3, 4)
	fmt.Println("Sum Total is :", total)

	fmt.Println("======== Recursive functions  =============")
	factValue := fact(5)

	fmt.Println("Factorial Value is :", factValue)

	fmt.Println("======== Struct in Go  =============")

	type Person struct {
		name string
		age  int
	}

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	person := Person{name: "Messi", age: 34}

	fmt.Println("The name of person is ", person.name)

	fmt.Println("Struct with default values", dog)

	fmt.Println("======== Methods on Struct  =============")

	r := rect{width: 10, height: 5}

	fmt.Println("Area of rectangle is ", r.area())

	fmt.Println("======== Interface in Go  =============") //Interfaces are named collections of method signatures.
	// p := rect{width: 3, height: 4}

	// measure(p)

}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type geometry interface {
	area() int
	perim() int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func addData(a, b int) int {
	return a + b
}

func returnMulipelVal(x, y int) (int, int) {
	return x, y
}

func sums(nums ...int) int {
	fmt.Print(nums, "	")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total

}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
