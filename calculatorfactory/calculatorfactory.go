package calculatorfactory

import (
	"github.com/cclulu/practiceCLI/calculator"
	"github.com/mitchellh/cli"
)

func Calculatorfactory() (cli.Command, error) {
	return &calculator.CalculatorCommand{}, nil
}
