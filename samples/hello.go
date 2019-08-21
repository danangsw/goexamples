package samples

import "fmt"

// Hello This func must be Exported, Capitalized, and comment added.
func Hello(name string) {
	if len(name) == 0 { name = "Dolly" }
	fmt.Printf("Hello %s\n", name)
}