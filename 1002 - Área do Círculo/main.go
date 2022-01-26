package main

import "fmt"

func main() {
	const pi = 3.14159
	var raio float64

	fmt.Scanf("%f\n", &raio)

	area := pi * raio * raio

	fmt.Printf("A=%.4f\n", area)
}
