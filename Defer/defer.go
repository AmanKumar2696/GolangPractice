package main

import "fmt"

func main() {
	defer fmt.Println("Deffered Function")
	fmt.Println("2nd Function")
}
