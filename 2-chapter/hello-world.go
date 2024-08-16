package main

import "fmt"


// Para fazer declarações fora de um code block precisa usar var
var scope = "I have the high ground"

func main() {
	_, errors := fmt.Println("Hello World!") // Hello World!
	// Podemos usar "_" para ignorar um valor retornado na função
	fmt.Println(errors) // <nil>

	// Tipos de variáveis
	x := 1 // Declaração :=
	y := "ce tá doido"
	z := true
	a := 10.0
	
	fmt.Println(x, y, z, a)
	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n", y, y)
	fmt.Printf("a: %v, %T \n", a, a)

	x = 12 // Atribuição =
	fmt.Printf("x: %v, %T \n", x, x)
}