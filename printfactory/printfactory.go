package printfactory

import (
	"github.com/cclulu/practiceCLI/printer"
	"github.com/mitchellh/cli"
)

func PrintFactory() (cli.Command, error) {
	return &printer.PrinterCommand{}, nil
}
