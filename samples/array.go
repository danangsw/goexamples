package samples

import (
	"fmt"
)

func Array() {
	arr := [10]int{1, 2, 3, 4, 5}
	fmt.Println("Len: ", len(arr))
	fmt.Println("Array: ", arr)

	// Accessing elements
	fmt.Println("First element: ", arr[0])
	fmt.Println("Second element: ", arr[1])
	fmt.Println("Third element: ", arr[2])

	//Slicing
	fmt.Println("Slice from index 1 to 4: ", arr[1:4])

	//Set
	arr[9] = 10
	fmt.Println("Updated Array: ", arr)

	//Two Dimensional Array
	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Two Dimensional Array: ", twoD)
	fmt.Println("Two Dimensional Array Row 1: ", twoD[0])
	fmt.Println("Two Dimensional Array Row 2: ", twoD[1])

	//Update wo dimensional row 2
	twoD[1] = [3]int{7, 8, 9}
	fmt.Println("Update Two Dimensional Array Row 2: ", twoD)
}
