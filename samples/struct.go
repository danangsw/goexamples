package samples

import (
	"fmt"
)

type address struct {
	address string
	phone   string
}

type student struct {
	name  string
	age   int
	grade int
	address
}

func StructExample() {
	s1 := student{
		name:  "John",
		age:   20,
		grade: 3,
	}
	s2 := student{"Jane", 22, 4, address{"Jakarta", "62888112134"}}
	s3 := student{name: "Gerry"}
	s4 := student{grade: 5}
	s5 := student{age: 25, grade: 2, name: "Gina"}

	// Print all student details
	printStudentDetails(s1)
	printStudentDetails(s2)
	printStudentDetails(s3)
	printStudentDetails(s4)
	printStudentDetails(s5)

	// Pointer struct
	var s6 *student = &s1
	s6.grade = 4
	s1.age = 21

	a1 := address{"Bandung", "628546713557"}
	s1.address = a1
	fmt.Println("After struct pointer variable definition")
	printStudentDetails(s1)
	printStudentDetails(*s6)

	// Anonymous struct
	s7 := struct {
		name  string
		age   int
		grade int
	}{
		"Agus", 27, 5,
	}

	s8 := student{s7.name, s7.age, s7.grade, address{}}

	fmt.Println("All Students")
	students := []student{}
	students = append(students, s1, s2, s3, s4, s5, *s6, s8)
	printAllStudentDetails(students)
}

func printAllStudentDetails(students []student) {
	for _, student := range students {
		printStudentDetails(student)
	}
}

func printStudentDetails(s student) {
	fmt.Printf("Student: %v. \n", &s)
	fmt.Printf("Name: %s, Age: %d, Grade: %d, Address: %s, Phone: %s.\n", s.name, s.age, s.grade, s.address.address, s.address.phone)
}
