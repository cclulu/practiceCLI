package printer

import (
	"fmt"

	"github.com/mitchellh/cli"
)

var _ cli.Command = (*PrinterCommand)(nil)

type PrinterCommand struct{}

func (p PrinterCommand) Help() string {
	helpText := ` 
Uses: 

print <strings>

This will just print everything that you give it
`
	return helpText
}

func (p PrinterCommand) Run(args []string) int {
	return p.run(args)
}

func (p PrinterCommand) Synopsis() string {
	return "print text on screen"
}

func (p PrinterCommand) run(args []string) int {
	if len(args) == 0 {
		println("Nothing was provided")
	}

	var stringsToPrint string

	for _, arg := range args {
		stringsToPrint += fmt.Sprintf("%s ", arg)
	}

	fmt.Println(stringsToPrint)

	return 0
}
