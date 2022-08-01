package mkdir

import (
	"fmt"
)

type Cmd struct{}

func (_ Cmd) Execute(args []string) error {
	fmt.Println("test")
	fmt.Println("test")
	return nil
}
