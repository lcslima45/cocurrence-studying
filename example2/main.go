package main

import "fmt"

//esse programa da erro panic pois não há goroutines interagindo com o channels fora do main
func main() {
	var ch chan int
	ch = make(chan int)
	ch <- 10
	v := <-ch
	fmt.Println("received", v)
}
