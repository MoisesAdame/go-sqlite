// Author: Moisés Adame-Aguilar
// Creation Date: March 20th, 2024

package commands

import (
	"../buffer"
	"fmt"
	"os"
)

// MetaCommandResult represents a command result.
type MetaCommandResult struct {
    IsSuccesful bool
}

// NewUnsuccesfulCommand returns an instance of a successful MetaCommandResult.
func NewUnsuccesfulCommand() MetaCommandResult {
	return MetaCommandResult{false}
}

// NewUnsuccesfulCommand returns an instance of an unsuccessful MetaCommandResult.
func NewSuccessfulCommand() MetaCommandResult {
	return MetaCommandResult{true}
}

// ExecuteMetaCommand checks for a command and executes it.
func ExecuteMetaCommand(ib *buffer.InputBuffer) MetaCommandResult {
	if(ib.ToString() == ".exit") {
		fmt.Println("Bye!")
		os.Exit(1)
		return NewSuccessfulCommand()
	}else{
		return NewUnsuccesfulCommand()
	}
}
