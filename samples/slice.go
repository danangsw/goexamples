package samples

import (
	"fmt"
)

func SliceExample() {
	var fruits = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Printf("Fruits : %v , Length : %d, Capacity : %d\n", fruits, len(fruits), cap(fruits))

	var newFruits = fruits[1:3]
	fmt.Printf("New Fruits : %v , Length : %d, Capacity : %d\n", newFruits, len(newFruits), cap(newFruits))
	newFruits[0] = "Blueberry"
	fmt.Printf("New Fruits After Update : %v , Length : %d, Capacity : %d\n", newFruits, len(newFruits), cap(newFruits))
	fmt.Printf("Fruits, After New Fruits Update : %v , Length : %d, Capacity : %d\n", fruits, len(fruits), cap(fruits))

	var newFruits2 = newFruits[1:3]
	fmt.Printf("New Fruits 2 : %v , Length : %d, Capacity : %d\n", newFruits2, len(newFruits2), cap(newFruits2))
	newFruits2 = append(newFruits2, "Pinnaple")
	fmt.Printf("New Fruits 2 Append Pinnaple : %v , Length : %d, Capacity : %d\n", newFruits2, len(newFruits2), cap(newFruits2))
	newFruits2 = append(newFruits2, "Grape")
	fmt.Printf("New Fruits 2 Append Grape : %v , Length : %d, Capacity : %d\n", newFruits2, len(newFruits2), cap(newFruits2))

	// Copy fruits
	var otherFruits = make([]string, len(fruits)-2)
	var n = copy(otherFruits, fruits)
	fmt.Printf("Other Fruits, copy %d items : %v , Length : %d, Capacity : %d\n", n, otherFruits, len(otherFruits), cap(otherFruits))
	otherFruits[0] = "Avocado"
	fmt.Printf("Other Fruits After Update : %v , Length : %d, Capacity : %d\n", otherFruits, len(otherFruits), cap(otherFruits))
	fmt.Printf("Fruits, After Other Fruits Update : %v , Length : %d, Capacity : %d\n", fruits, len(fruits), cap(fruits))

	// 3 Index slicing
	var newFruits3 = fruits[0:2:2]
	fmt.Printf("New Fruits 3: %v , Length : %d, Capacity : %d\n", newFruits3, len(newFruits3), cap(newFruits3))
	var newFruits4 = fruits[0:2]
	fmt.Printf("New Fruits 4: %v , Length : %d, Capacity : %d\n", newFruits4, len(newFruits4), cap(newFruits4))
}
