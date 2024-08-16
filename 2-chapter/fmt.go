// " " para interpreted string literals
//  `` para raw string literals

package main

import (
	"fmt"
)

func main() {
	// Prints
	// fmt.Print("Opa")
	// fmt.Printf("Opa")
	x := "Oi"
	y := "Bom dia"

	fmt.Println(x)
	fmt.Println(y)

	// SPrints
	z := fmt.Sprint(x, ", ", y)

	fmt.Println(z)

	// FPrints - para passar informações para arquivos ou entre conexões com db e http

}