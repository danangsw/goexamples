package samples

import (
	"fmt"
	"math"
)

func ForSamples() {
	fmt.Println("For Loop 1:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("For Loop 2:")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("For Loop Range:")
	for i := 0; i < 3; i++ {
		fmt.Println("Range: ", i)
	}

	fmt.Println("For Loop Single:")
	for {
		fmt.Println("Loop")
		break
	}

	fmt.Println("For Loop Conditional:")
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Odd: ", n)
	}

	fmt.Println("For continued:")
	var sum int = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum:", sum)

	fmt.Println("For is GO's 'while':")
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum:", sum)

	fmt.Println("Exercise: Newton's method")
	x := float64(1.2)
	fmt.Printf("Sqrt(%f): %f\n", x, sqrt(x))
	fmt.Printf("Math.Sqrt(%f): %f\n", x, math.Sqrt(x))
}

func sqrt(x float64) float64 {
	y := float64(1)
	z := float64(1)
	lv := 0.00000001
	it := 20
	for i := 0; i < it; i++ {
		y = z
		z -= (z*z - x) / (2 * z)
		if y == z || z < lv {
			break
		}
	}

	return z
}
