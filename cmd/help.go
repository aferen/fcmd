package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type HelpCmd struct{}

func (_ HelpCmd) Execute(args []string) error {
	return printUsage()
}

func printUsage() error {
	help := `Usage: fcmd [COMMAND] [OPTION]

Commands:
  mkdir

Use "fcmd <command> --help/-h" for more information about a given command.`

	_, err := fmt.Fprintf(os.Stdout, "%s\n", help)
	return errors.Wrap(err, "write error")
}
