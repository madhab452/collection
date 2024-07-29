package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! !\n", user.Username)
	fmt.Printf("feel free to type any statements\n")
	Start(os.Stdin, os.Stdout)
}
