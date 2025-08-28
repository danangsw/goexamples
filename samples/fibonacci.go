package samples

//import "fmt"

// Fibonacci Sequence is the series of numbers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
func Fibonacci(number int) int  {
	if number <= 1 {
		//fmt.Println("base case => ", number);
		return number
	}

	//fmt.Println("number - 2 => ", number - 2, "number - 1 => ", number - 1);
	return Fibonacci(number - 2) + Fibonacci(number - 1)
}