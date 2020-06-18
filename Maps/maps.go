package main

import "fmt"

func main() {
	prodPrices := make(map[string]int)
	prodPrices["Mouse"] = 600
	prodPrices["Keyboard"] = 1000
	prodPrices["Monitor"] = 5500
	fmt.Println(prodPrices, "nIteration")

	for key, val := range prodPrices {
		fmt.Println("Product: ", key, "Price: ", val)
	}
	prodPrices["Printer"] = 8250
	fmt.Println("Modifying Map:", prodPrices)
	delete(prodPrices, "Mouse")
	fmt.Println(prodPrices)

}
