package samples

import "fmt"

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
}
