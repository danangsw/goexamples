package samples

import (
	"fmt"
	"os"
	"sync"
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
	fmt.Println("1. File processing")
	err := ProcessFile("./samples/defer_test.txt")
	if err != nil {
		fmt.Println("Error processing file:", err)
	}

	// 2. Mutex unlocking example
	fmt.Println("2. Mutex unlocking example")
	counter := SafeCounter{}
	for i := 0; i < 10; i++ {
		counter.Increment()
	}
	fmt.Println("Counter:", counter.GetCount())

	// 3. Recovering from panic
	fmt.Println("3. Recovering from panic")
	recoverFromPanic()
}

// 1. File processing
func ProcessFile(filename string) error {
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
type SafeCounter struct {
	count int
	mu    sync.Mutex // Removed sync.Mutex to avoid import and concurrent
}

func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         // Lock the mutex before incrementing
	defer c.mu.Unlock() // Endure mutex is unlocked when function returns
	c.count++
}

// 3. Recovering from panics
func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate panic
	panic("Something went wrong!")
}
