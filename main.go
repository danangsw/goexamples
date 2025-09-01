package main

import (
	"flag"
	"fmt"
	"os"

	samples "./samples"
)

const slice = "slice"
const array = "array"
const switches = "switch"
const switchTimeFlag = "time"
const switchTypeFlag = "type"
const branching = "branching"
const branchingValueFlag = "value"
const hello = "hello"
const helloNameFlag = "name"
const value = "values"
const variable = "variables"
const exit = "exit"
const reverse = "reverse"
const reverseStringFlag = "string"
const fibonacci = "fibonacci"
const fibonacciNumberFlag = "number"
const nthprime = "nthprime"
const nthprimeNumberFlag = "number"
const nthprimeisPrintedFlag = "isprinted"
const defaultMsg = "let's go! learn golang by examples."

func main() {
	sliceCmd := flag.NewFlagSet(slice, flag.ExitOnError)

	arrayCmd := flag.NewFlagSet(array, flag.ExitOnError)

	switchCmd := flag.NewFlagSet(switches, flag.ExitOnError)
	switchDay := switchCmd.String(switchTimeFlag, "", "switch -time=<day | weekday | time>")
	switchType := switchCmd.String(switchTypeFlag, "", "switch -type=<Your type here>")

	helloCmd := flag.NewFlagSet(hello, flag.ExitOnError)
	helloName := helloCmd.String(helloNameFlag, "", "hello -name=<Your name here>")

	branchingCmd := flag.NewFlagSet(branching, flag.ExitOnError)
	branchingValue := branchingCmd.String(branchingValueFlag, "", "branching -value=<Your value here>")

	valueCmd := flag.NewFlagSet(value, flag.ExitOnError)

	variableCmd := flag.NewFlagSet(variable, flag.ExitOnError)

	exitCmd := flag.NewFlagSet(exit, flag.ExitOnError)

	reverseCmd := flag.NewFlagSet(reverse, flag.ExitOnError)
	reverseString := reverseCmd.String(reverseStringFlag, "", "reverse -string=<Your text here>")

	fibonacciCmd := flag.NewFlagSet(fibonacci, flag.ExitOnError)
	fibonacciNumber := fibonacciCmd.Int(fibonacciNumberFlag, 0, "fibonacci -number=<Your number here>")

	nthprimeCmd := flag.NewFlagSet(nthprime, flag.ExitOnError)
	nthPrimeNumber := nthprimeCmd.Int(nthprimeNumberFlag, 0, "nthprime -number=<Your number here>")
	nthPrimeIsPrinted := nthprimeCmd.Bool(nthprimeisPrintedFlag, false, "nthprime -isprinted=<true | false>")

	if len(os.Args) < 2 {
		fmt.Println(defaultMsg)
		os.Exit(1)
	}

	switch os.Args[1] {
	case slice:
		sliceCmd.Parse(os.Args[2:])
		samples.SliceExample()
	case array:
		arrayCmd.Parse(os.Args[2:])
		samples.Array()
	case switches:
		switchCmd.Parse(os.Args[2:])
		samples.Switches(*switchDay, *switchType)
	case branching:
		branchingCmd.Parse(os.Args[2:])
		samples.Branching(*branchingValue)
	case hello:
		helloCmd.Parse(os.Args[2:])
		samples.Hello(*helloName)
	case value:
		valueCmd.Parse(os.Args[2:])
		samples.Values()
	case variable:
		variableCmd.Parse(os.Args[2:])
		samples.Variables()
	case reverse:
		reverseCmd.Parse(os.Args[2:])
		result := samples.Reverse(*reverseString)
		fmt.Println(result)
	case fibonacci:
		fibonacciCmd.Parse(os.Args[2:])
		result := samples.Fibonacci(*fibonacciNumber)
		fmt.Println(result)
	case nthprime:
		nthprimeCmd.Parse(os.Args[2:])
		result := samples.NthPrime(*nthPrimeNumber, *nthPrimeIsPrinted)
		fmt.Println(result)
	case exit:
		exitCmd.Parse(os.Args[2:])
		samples.Exit()
	default:
		fmt.Println(defaultMsg)
		os.Exit(2)
	}
}
