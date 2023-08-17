package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mattbidewell/gloop/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, This is the Gloop programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
