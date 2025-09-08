package samples

import "fmt"

func DeferSamples() {
	defer fmt.Println("world")

	// Deferred calls are executed in last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}
