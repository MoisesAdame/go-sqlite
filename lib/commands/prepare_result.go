// Author: Moisés Adame-Aguilar
// Creation Date: March 20th, 2024

package commands

import (
	"../buffer"
	"fmt"
)

// PrepareResult represents a command result.
type PrepareResult struct {
    IsSuccesful bool
}

// NewUnsuccesfulCommand returns an instance of a successful PrepareResult.
func NewUnsuccesfulPrepare() PrepareResult {
	return PrepareResult{false}
}

// NewUnsuccesfulCommand returns an instance of an unsuccessful PrepareResults.
func NewSuccessfulPrepare() PrepareResult {
	return PrepareResult{true}
}

// Statement represents a command result.
type Statement struct {
    Type string
}

// NewStatement returns a pointer to a Statement whose type is null.
func NewStatement() *Statement {
	return &Statement{}
}

// NewInsertStatement returns a pointer to a Statement whose type is select.
func NewSelectStatement() *Statement {
	return &Statement{"select"}
}

// NewInsertStatement returns a pointer to a Statement whose type is insert.
func NewInsertStatement() *Statement {
	return &Statement{"insert"}
}

// PrepareStatement based on the command input given, returns a PrepareResult.
func PrepareStatement(ib *buffer.InputBuffer, statement *Statement) PrepareResult {

	bufferString := ib.ToString()

	if(len(bufferString) >= 6 && bufferString[:6] == "insert"){
		statement.Type = "insert"
		return NewSuccessfulPrepare()
	}else if(bufferString == "select"){
		statement.Type = "select"
		return NewSuccessfulPrepare()
	}else{
		return NewUnsuccesfulPrepare()
	}
}

// ExecuteStatement prints the execution of a given command, based on the statement type.
func ExecuteStatement(statement *Statement) {
	if(statement.Type == "select"){
		fmt.Println("Executing a SELECT.")
	}else if(statement.Type == "insert") {
		fmt.Println("Executing an INSERT.")
	}
}
