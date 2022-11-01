package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeProduct := nikeFactory.makeProduct()
	adidasProduct := adidasFactory.makeProduct()

	printProductDetails(nikeProduct)
	printProductDetails(adidasProduct)
}

func printProductDetails(s IProduct) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
