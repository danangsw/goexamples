package samples

import "fmt"

// Hello is our first program will print the classic “hello world” message.
func Hello(name string) {
	if len(name) == 0 { name = "Dolly" }
	fmt.Printf("Hello %s\n", name)
}