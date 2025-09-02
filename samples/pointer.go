package samples

import (
	"fmt"
)

func PointerExample() {
	var number int = 77
	var name string = "Gogon"

	var ptrNumber *int = &number
	var ptrName *string = &name

	fmt.Println("Pointer Example:")
	fmt.Printf("Address / Value of number: %v / %v \n", &number, number)
	fmt.Printf("Address / Value of name: %v / %v \n", &name, name)
	fmt.Printf("Address / Value of ptrNumber: %v / %v \n", ptrNumber, *ptrNumber)
	fmt.Printf("Address / Value of ptrName: %v / %v \n", ptrName, *ptrName)

	// Change the value of the original variables using pointers
	*ptrNumber = 100
	*ptrName = "Gogon Updated"

	fmt.Println("After updating pointer values:")
	fmt.Printf("Address / Value of number: %v / %v \n", &number, number)
	fmt.Printf("Address / Value of name: %v / %v \n", &name, name)
	fmt.Printf("Address / Value of ptrNumber: %v / %v \n", ptrNumber, *ptrNumber)
	fmt.Printf("Address / Value of ptrName: %v / %v \n", ptrName, *ptrName)

	// Change the original value variables
	number = 1000
	name = "Gogon Final Updated"

	fmt.Println("After updating original values:")
	fmt.Printf("Address / Value of number: %v / %v \n", &number, number)
	fmt.Printf("Address / Value of name: %v / %v \n", &name, name)
	fmt.Printf("Address / Value of ptrNumber: %v / %v \n", ptrNumber, *ptrNumber)
	fmt.Printf("Address / Value of ptrName: %v / %v \n", ptrName, *ptrName)

	// Swap variables
	var swapNumber int = 1
	var swapName = "John"

	swapInt(&swapNumber, ptrNumber)
	swapString(&swapName, ptrName)

	fmt.Println("After swap values:")
	fmt.Printf("Address / Value of number: %v / %v \n", &number, number)
	fmt.Printf("Address / Value of name: %v / %v \n", &name, name)
	fmt.Printf("Address / Value of ptrNumber: %v / %v \n", ptrNumber, *ptrNumber)
	fmt.Printf("Address / Value of ptrName: %v / %v \n", ptrName, *ptrName)
	fmt.Printf("Address / Value of swapNumber: %v / %v \n", &swapNumber, swapNumber)
	fmt.Printf("Address / Value of swapName: %v / %v \n", &swapName, swapName)
}

func swapInt(a, b *int) {
	*a, *b = *b, *a
}

func swapString(a, b *string) {
	*a, *b = *b, *a
}
