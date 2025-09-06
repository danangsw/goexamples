package samples

import (
	"fmt"
	"math"
	"time"
)

const Pi = 3.14159
const Greeting = "Hello, Go!"
const n = 5000000
const d = 3e20 / 5
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int {
	return x*10 + 1
}

func NeedFloat(x float64) float64 {
	return x*0.1 + 1
}

// Custom type based on a constant
type FileSize int64

const (
	KB FileSize = 1 << 10 // 1024
	MB FileSize = 1 << 20 // 1024*1024 : 1048576
	GB FileSize = 1 << 30 // 1024*1024*1024 : 1073741824
	TB FileSize = 1 << 40 // 1024*1024*1024*1024 : 1099511627776
)

func (fs FileSize) String() string {
	switch {
	case fs >= TB:
		return fmt.Sprintf("%.2f TB", float64(fs)/float64(TB))
	case fs >= GB:
		return fmt.Sprintf("%.2f GB", float64(fs)/float64(GB))
	case fs >= MB:
		return fmt.Sprintf("%.2f MB", float64(fs)/float64(MB))
	case fs >= KB:
		return fmt.Sprintf("%.2f KB", float64(fs)/float64(KB))
	default:
		return fmt.Sprintf("%d B", fs)
	}
}

// Using constants for bitwise operations (flags)
const (
	FlagReadable   = 1 << iota // 1
	FlagWritable               // 2
	FlagExecutable             // 4
)

func checkPermissions(flags int) {
	fmt.Println("Readable:", flags&FlagReadable != 0)
	fmt.Println("Writable:", flags&FlagWritable != 0)
	fmt.Println("Executable:", flags&FlagExecutable != 0)
}

// Constant for defining timeouts
const (
	ShortTimeout  = 500 * time.Millisecond
	MediumTimeout = 2 * time.Second
	LongTimeout   = 10 * time.Second
)

func performOperationWithTimeout(timeout time.Duration) {
	fmt.Println("Performing operation with timeout:", timeout)
	time.Sleep(timeout)
	fmt.Println("Operation completed (or timeout)")
}

const (
	ApiVersion = "v1"
	ApiBaseUrl = "https://api.example.com/" + ApiVersion
)

func getApiEndpoint(resource string) string {
	return ApiBaseUrl + "/" + resource
}

func ConstantSamples() {
	// Basic used
	fmt.Println("Pi:", Pi)
	fmt.Println(Greeting)
	fmt.Println("n:", n)
	fmt.Println("d:", d)

	// Performing calculations
	fmt.Println("Square root of n: ", math.Sqrt(n))
	fmt.Println("Sin of n:", math.Sin((n)))
	fmt.Println(Pi)
	fmt.Println(NeedInt(Small))
	fmt.Println(NeedFloat(Small))
	fmt.Println(NeedFloat(Big))

	// Custom file size examples
	var fileSize FileSize = 2684354560 // 2.5 GB
	fmt.Println("FileSize:", fileSize)

	// Permissions examples
	checkPermissions(FlagReadable | FlagWritable | FlagExecutable)
	checkPermissions(FlagReadable | FlagWritable)
	checkPermissions(FlagExecutable)

	// Timeout examples
	performOperationWithTimeout(ShortTimeout)
	performOperationWithTimeout(MediumTimeout)
	performOperationWithTimeout(LongTimeout)

	// API endpoint examples
	fmt.Println("API endpoint for users:", getApiEndpoint("users"))
}
