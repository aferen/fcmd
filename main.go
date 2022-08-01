package main

import (
	"errors"
	"fcmd/cmd/mkdir"
	"log"
	"os"
)

type Cmd interface {
	Execute(args []string) error
}

func main() {
	cmd, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute(os.Args[2:])
}

func parseArgs(argv []string) (Cmd, error) {
	if len(argv) == 1 {
		cmd := argv[0]
		if cmd == "mkdir" {
			return mkdir.Cmd{}, nil
		}
	}
	return nil, errors.New("invalid Command")
}
