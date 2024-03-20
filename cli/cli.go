// Author: Moisés Adame-Aguilar
// Creation Date: March 18th, 2024

package cli

import (
	"os"
	"../lib/models"
	"fmt"
)

type CLI struct {}

// Constructor method for the Command Line Interface.
func NewCLI() *CLI {
	return &CLI{}
}

// readInput prints the classic sqlite prompt.
func printPrompt() {
	fmt.Printf("db > ")
}

// readInput reads from the user input and stores it as a buffer.
func readInput(ib *lib.InputBuffer) {
	var input string
	fmt.Scan(&input)
	ib.ReadFromString(input)
}

func (cli *CLI) Run(){

	inputBuffer := lib.NewInputBuffer()

	for {
		printPrompt()

		readInput(inputBuffer)

		if(inputBuffer.ToString() == ".exit") {
			fmt.Println("Bye!")
			os.Exit(1)
		}else{
			fmt.Println("Unrecognized command.")
		}
    }
}
