// Author: Moisés Adame-Aguilar
// Creation Date: March 18th, 2024

package cli

import (
	"../buffer"
	"../commands"
	"fmt"
	"bufio"
	"os"
	"../models"
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
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

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
		}else{
			table := models.NewTable(3)

			statement := commands.NewStatement()
			prepare := commands.PrepareStatement(inputBuffer, statement, table)

			if(prepare.IsSuccesful){
				commands.ExecuteStatement(statement)
			}else{
				fmt.Println("Unrecognized keyword at start of " + inputBuffer.ToString())
			}
		}
    }
}
