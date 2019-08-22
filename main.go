package main

import (
	"flag"
	"fmt"
	"os"
	samples "github.com/danangsw/goexamples/samples"
)

const hello = "hello"
const helloNameFlag = "name"
const value = "values"
const variable = "variables"
const exit = "exit"
const reverse = "reverse"
const reverseStringFlag = "string"
const defaultMsg = "let's go! learn golang by examples."

func main () {
	helloCmd := flag.NewFlagSet(hello, flag.ExitOnError)
	helloName := helloCmd.String(helloNameFlag, "", "hello -name=<Your name here>")

	valueCmd := flag.NewFlagSet(value, flag.ExitOnError)

	variableCmd := flag.NewFlagSet(variable, flag.ExitOnError)

	exitCmd := flag.NewFlagSet(exit, flag.ExitOnError)

	reverseCmd := flag.NewFlagSet(reverse, flag.ExitOnError)
	reverseString := reverseCmd.String(reverseStringFlag, "", "reverse -string=<Your text here>")

	if len(os.Args) < 2 {
		fmt.Println(defaultMsg)
		os.Exit(1)	
	}

	switch os.Args[1] {
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
		case exit:
			exitCmd.Parse(os.Args[2:])
			samples.Exit()
		default:
			fmt.Println(defaultMsg)
			os.Exit(2)
	}
}