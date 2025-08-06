package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/santos-404/myte/repl"
)

const ASCII_ART = `░▒▓██████████████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░   ░▒▓█▓▒░   ░▒▓██████▓▒░   
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░      ░▒▓█▓▒░   ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░      ░▒▓█▓▒░   ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░      ░▒▓█▓▒░   ░▒▓████████▓▒░ 
                                                                `

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println(ASCII_ART)
	fmt.Printf("Hi there %s!, this is the Myte programming language!\n", 
		user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

