package main

import "fmt"

func main() {
	go hello()
}

func hello() {
	fmt.Println("Hi :D")
}
