package samples

import "fmt"

func MapExample() {
	person := map[string]int{
		"John":  30,
		"Jane":  25,
		"Bob":   35,
		"Alice": 28,
	}

	for name, age := range person {
		fmt.Printf("%s is %d years old. \n", name, age)
	}
	fmt.Printf("# Length of Person : %d \n", len(person))

	// Add items
	person["Tom"] = 40
	person["Charlie"] = 51
	fmt.Println("Person, After adding Tom and Charlie :", person)
	fmt.Printf("# Length of Person : %d \n", len(person))

	// Update items
	person["Alice"] = 29
	fmt.Println("Person, After updating Alice's age :", person)
	fmt.Printf("# Length of Person : %d \n", len(person))

	// Find items
	find := "Bob"
	value, isExist := person[find]
	if isExist {
		fmt.Printf("%v is %d years old. \n", find, value)
	} else {
		fmt.Printf("%v is not found. \n", find)
	}

	find = "Alex"
	value, isExist = person[find]
	if isExist {
		fmt.Printf("%v is %d years old. \n", find, value)
	} else {
		fmt.Printf("%v is not found. \n", find)
	}

	// Delete items
	delete(person, "John")
	delete(person, "Jane")
	fmt.Println("Person, After deleting John and Jane :", person)
	fmt.Printf("# Length of Person : %d \n", len(person))

	delete(person, "Alex") // No error even item is not found
	fmt.Println("Person, After deleting Alex :", person)

	// Map and Slice
	names := make([]string, 0, len(person))
	fmt.Printf("Properties of names (before copy): len: %v, cap: %v, values: %v. \n", len(names), cap(names), names)

	for name := range person {
		names = append(names, name)
	}
	fmt.Printf("Properties of names (after copy): len %v, cap %v, values %v. \n", len(names), cap(names), names)
	names = append(names, "Tom", "Jerry", "Boby")
	fmt.Printf("Properties of names (after append): len %v, cap %v, values %v. \n", len(names), cap(names), names)
	for i := 0; i < len(names); i++ {
		if value, found := person[names[i]]; found {
			fmt.Printf("%v is %d years old. \n", names[i], value)
		} else {
			fmt.Printf("%v is not found. \n", names[i])
		}
	}

	var data = []map[string]string{
		{"name": "John", "gender": "male", "age": "30"},
		{"address": "123 Main St, Anytown, USA", "phone": "555-1234"},
		{"community": "Anytown", "citizen": "USA"},
	}

	fmt.Printf("Data properties : len %v, cap %v, values %v\n", len(data), cap(data), data)
	for _, item := range data {
		fmt.Println("Item :", item)
		for key, value := range item {
			fmt.Printf("%v : %v \n", key, value)
		}
	}
}
