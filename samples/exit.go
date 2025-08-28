package samples

import (
	"fmt"
	"os"
)

// Exit Use os.Exit to immediately exit with a given status.
func Exit()  {
	defer fmt.Println("!")

	os.Exit(3)
}