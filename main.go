package main

import (
	"flag"
	"fmt"
	"os"

	samples "./samples"
)

const constantFlag = "constant"
const basicTypesFlag = "types"
const listNodeFlag = "listnode"
const structFlag = "struct"
const pointerFlag = "pointer"
const funcFlag = "func"
const mapFlag = "map"
const sliceFlag = "slice"
const arrayFlag = "array"
const switchesFlag = "switch"
const switchTimeFlag = "time"
const switchTypeFlag = "type"
const branchingFlag = "branching"
const branchingValueFlag = "value"
const helloFlag = "hello"
const helloNameFlag = "name"
const valueFlag = "values"
const variableFlag = "variables"
const exitFlag = "exit"
const reverseFlag = "reverse"
const reverseStringFlag = "string"
const fibonacciFlag = "fibonacci"
const fibonacciNumberFlag = "number"
const nthprimeFlag = "nthprime"
const nthprimeNumberFlag = "number"
const nthprimeisPrintedFlag = "isprinted"
const defaultMsg = "let's go! learn golang by examples."

func main() {
	constantCmd := flag.NewFlagSet(constantFlag, flag.ExitOnError)

	basicTypesCmd := flag.NewFlagSet(basicTypesFlag, flag.ExitOnError)

	listNodeCmd := flag.NewFlagSet(listNodeFlag, flag.ExitOnError)

	structCmd := flag.NewFlagSet(structFlag, flag.ExitOnError)

	pointerCmd := flag.NewFlagSet(pointerFlag, flag.ExitOnError)

	funcCmd := flag.NewFlagSet(funcFlag, flag.ExitOnError)
	funcContain := funcCmd.String("c", "", "func -c=<Your string contain here>")
	funcPrefix := funcCmd.String("p", "", "func -p=<Your string prefix here>")
	funcLength := funcCmd.Int("l", 0, "func -l=<Your string min-length here>")

	mapCmd := flag.NewFlagSet(mapFlag, flag.ExitOnError)

	sliceCmd := flag.NewFlagSet(sliceFlag, flag.ExitOnError)

	arrayCmd := flag.NewFlagSet(arrayFlag, flag.ExitOnError)

	switchCmd := flag.NewFlagSet(switchesFlag, flag.ExitOnError)
	switchDay := switchCmd.String(switchTimeFlag, "", "switch -time=<day | weekday | time>")
	switchType := switchCmd.String(switchTypeFlag, "", "switch -type=<Your type here>")

	helloCmd := flag.NewFlagSet(helloFlag, flag.ExitOnError)
	helloName := helloCmd.String(helloNameFlag, "", "hello -name=<Your name here>")

	branchingCmd := flag.NewFlagSet(branchingFlag, flag.ExitOnError)
	branchingValue := branchingCmd.String(branchingValueFlag, "", "branching -value=<Your value here>")

	valueCmd := flag.NewFlagSet(valueFlag, flag.ExitOnError)

	variableCmd := flag.NewFlagSet(variableFlag, flag.ExitOnError)

	exitCmd := flag.NewFlagSet(exitFlag, flag.ExitOnError)

	reverseCmd := flag.NewFlagSet(reverseFlag, flag.ExitOnError)
	reverseString := reverseCmd.String(reverseStringFlag, "", "reverse -string=<Your text here>")

	fibonacciCmd := flag.NewFlagSet(fibonacciFlag, flag.ExitOnError)
	fibonacciNumber := fibonacciCmd.Int(fibonacciNumberFlag, 0, "fibonacci -number=<Your number here>")

	nthprimeCmd := flag.NewFlagSet(nthprimeFlag, flag.ExitOnError)
	nthPrimeNumber := nthprimeCmd.Int(nthprimeNumberFlag, 0, "nthprime -number=<Your number here>")
	nthPrimeIsPrinted := nthprimeCmd.Bool(nthprimeisPrintedFlag, false, "nthprime -isprinted=<true | false>")

	if len(os.Args) < 2 {
		fmt.Println(defaultMsg)
		os.Exit(1)
	}

	switch os.Args[1] {
	case constantFlag:
		constantCmd.Parse(os.Args[2:])
		samples.ConstantSamples()
	case basicTypesFlag:
		basicTypesCmd.Parse(os.Args[2:])
		samples.BasicTypesSamples()
	case listNodeFlag:
		listNodeCmd.Parse(os.Args[2:])
		samples.ListNodeExample()
	case structFlag:
		structCmd.Parse(os.Args[2:])
		samples.StructExample()
	case pointerFlag:
		pointerCmd.Parse(os.Args[2:])
		samples.PointerExample()
	case funcFlag:
		funcCmd.Parse(os.Args[2:])
		samples.FunctionExample(*funcContain, *funcPrefix, *funcLength)
	case mapFlag:
		mapCmd.Parse(os.Args[2:])
		samples.MapExample()
	case sliceFlag:
		sliceCmd.Parse(os.Args[2:])
		samples.SliceExample()
	case arrayFlag:
		arrayCmd.Parse(os.Args[2:])
		samples.Array()
	case switchesFlag:
		switchCmd.Parse(os.Args[2:])
		samples.Switches(*switchDay, *switchType)
	case branchingFlag:
		branchingCmd.Parse(os.Args[2:])
		samples.Branching(*branchingValue)
	case helloFlag:
		helloCmd.Parse(os.Args[2:])
		samples.Hello(*helloName)
	case valueFlag:
		valueCmd.Parse(os.Args[2:])
		samples.Values()
	case variableFlag:
		variableCmd.Parse(os.Args[2:])
		samples.Variables()
	case reverseFlag:
		reverseCmd.Parse(os.Args[2:])
		result := samples.Reverse(*reverseString)
		fmt.Println(result)
	case fibonacciFlag:
		fibonacciCmd.Parse(os.Args[2:])
		result := samples.Fibonacci(*fibonacciNumber)
		fmt.Println(result)
	case nthprimeFlag:
		nthprimeCmd.Parse(os.Args[2:])
		result := samples.NthPrime(*nthPrimeNumber, *nthPrimeIsPrinted)
		fmt.Println(result)
	case exitFlag:
		exitCmd.Parse(os.Args[2:])
		samples.Exit()
	default:
		fmt.Println(defaultMsg)
		os.Exit(2)
	}
}
