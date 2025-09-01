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

	//Update two dimensional row 2
	twoD[1] = [3]int{7, 8, 9}
	fmt.Println("Update Two Dimensional Array Row 2: ", twoD)

	// Three Dimensional Array
	threeD := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println("Three Dimensional Array: ", threeD)
	fmt.Println("Three Dimensional Array Row 1: ", threeD[0])
	fmt.Println("Three Dimensional Array Row 2: ", threeD[1])
	fmt.Println("Three Dimensional Array Row 3: ", threeD[2])

	// Make array
	var makeArray = make([]int, 5)
	fmt.Println("Make Array: ", makeArray)
	makeArray[0] = 1
	makeArray[2] = 2
	makeArray[4] = 5
	fmt.Println("Updated Make Array: ", makeArray)
	fmt.Println("Len of Make Array: ", len(makeArray))
	fmt.Println("Capacity of Make Array:", cap(makeArray))

	// Append
	makeArray = append(makeArray, 6)
	fmt.Println("Append 6 to Make Array:", makeArray)

	// [...] Array
	var numbers = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array with [...] :", numbers)
}
