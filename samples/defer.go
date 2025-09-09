package samples

import (
	"fmt"
	"os"
)

const (
	DeferSamplesBasic   = 1 << iota // 1
	DeferSampleAdvanced             // 2
)

func DeferSamples(types int) {
	switch types {
	case DeferSamplesBasic:
		deferSamplesBasic()
	case DeferSampleAdvanced:
		deferSamplesAdvanced()
	default:
		fmt.Printf("Undefined types, types = <%d | %d>\n", DeferSamplesBasic, DeferSampleAdvanced)
	}
}

// Deferred function samples basic
func deferSamplesBasic() {
	defer fmt.Println("world")

	// Deferred calls are executed in last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}

// Deferred function samples advanced
func deferSamplesAdvanced() {
	// 1. File processing example
	err := processFile("./samples/defer_test.txt")
	if err != nil {
		fmt.Println("Error processing file:", err)
	}
}

// 1. File processing
func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed when function returns

	// Read from the file
	buffer := make([]byte, 100)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Println("Read from file:", string(buffer))
	return nil
}

// 2. Mutex unlocking example

// 3. Recovering from panics
