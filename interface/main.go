package main

import "fmt"

// use use interfece to define method set
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	fmt.Println("================= Interface in Go =================")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// package main

// import "fmt"

// type englishBot struct{}

// type spanishBot struct{}

// // use use interfece to define method set
// type bot interface {
// 	getGreeting() string
// }

// func main() {

// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)

// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// // func printGreeting(eb englishBot) {
// // 	fmt.Println(eb.getGreeting())
// // }

// // func printGreeting(sb spanishBot) {
// // 	fmt.Println(sb.getGreeting())
// // }

// func (englishBot) getGreeting() string {

// 	return "Hello"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola"
// }
