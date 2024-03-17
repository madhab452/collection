package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/madhab452/collection/filter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! !\n", user.Username)
	fmt.Printf("feel free to type any statements\n")
	repl.Start(os.Stdin, os.Stdout)
}
