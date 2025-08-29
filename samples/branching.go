package samples

import (
	"fmt"
	"strconv"
)

func Branching(value string) {

	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
