package main

import (
	"fmt"
)

var a int
var b float32
var c string
var d bool

func main() {
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", b, b)
	fmt.Printf("%v, %T", c, c)
	fmt.Printf("%v, %T", d, d)
}