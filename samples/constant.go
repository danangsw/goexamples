package samples

import (
	"fmt"
	"math"
)

const Pi = 3.14159
const Greeting = "Hello, Go!"
const n = 5000000
const d = 3e20 / 5

func Constant() {
	fmt.Println("Pi:", Pi)
	fmt.Println(Greeting)
	fmt.Println("n:", n)
	fmt.Println("d:", d)
	// Performing calculations
	fmt.Println("Square root of n: ", math.Sqrt(n))
	fmt.Println("Sin of n:", math.Sin((n)))
}
