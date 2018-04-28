package main

import (
	"log"
	"os"

	"github.com/cclulu/practiceCLI/calculatorfactory"
	"github.com/cclulu/practiceCLI/printfactory"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("meetup", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"print": printfactory.PrintFactory,
		"calc":  calculatorfactory.Calculatorfactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
