package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mattbidewell/gloop/repl"
)

const logo = `
  .-"""-.
 /  ·  ·  \
(    ~     )
 ` + "`" + `'-----'` + "`" + `
   GLOOP
`

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Print(logo)
	fmt.Printf("Hello %s! Welcome to the Gloop programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
