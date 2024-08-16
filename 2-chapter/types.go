package main
import (
	"fmt"
)

type hotdog int
var b hotdog = 10

func main() {
	x := 10

	fmt.Printf("%v %T ", b, b)
	fmt.Println("")
	fmt.Printf("%v %T ", x, x)

	// b = x // cannot use x (variable of type int) as hotdog value in assignment

	// Conversão
	x = int(b)
}
