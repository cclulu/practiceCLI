package calculator

import (
	"github.com/mitchellh/cli"
)

var _ cli.Command = (*CalculatorCommand)(nil)

type CalculatorCommand struct{}

func (c CalculatorCommand) Help() string {
	helpText := ` 
Uses: 

calc <args> <number> <number>

This will preform the operation that you give it

options are: 

add
sub
mult

`
	return helpText
}

func (c CalculatorCommand) Run(args []string) int {
	return c.run(args)
}

func (c CalculatorCommand) Synopsis() string {
	return "arithmetic performed on the next two numbers"
}

func (c CalculatorCommand) run(args []string) int {
	// if len(args) == 0 {
	// 	println("Nothing was provided")
	// }
	// if len(args == 1) || len(args == 2) {
	// 	println("You need at least 3")
	// }

	// switch args[0] {
	// case strings.EqualFold("add"):
	// 	break
	// case strings.EqualFold("sub"):
	// 	break
	// case strings.EqualFold("mult"):
	// 	break
	// default:
	// 	println("Was not given a proper arguement")
	// }

	return 0
}
