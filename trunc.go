package main

import "fmt"

func main() {
	var number float64
	fmt.Println("Enter a number: ")
	fmt.Scanf("%f", &number)

	truncatedNumber := int(number)
	fmt.Printf("Truncated integer: %d\n", truncatedNumber)
}