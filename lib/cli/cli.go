// Author: Moisés Adame-Aguilar
// Creation Date: March 18th, 2024

package cli

import (
	"../buffer"
	"../commands"
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
func readInput(ib *buffer.InputBuffer) {
	var input string
	fmt.Scan(&input)
	ib.ReadFromString(input)
}

// Run prints prompt and receives input from user.
func (cli *CLI) Run(){

	inputBuffer := buffer.NewInputBuffer()

	for {
		printPrompt()

		readInput(inputBuffer)

		if(inputBuffer.ToString()[0] == byte('.')){
			command := commands.ExecuteMetaCommand(inputBuffer)

			if(!command.IsSuccesful) {
				fmt.Println("Unrecognized command " + inputBuffer.ToString())
			}
		}
    }
}
