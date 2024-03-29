// Author: Moisés Adame-Aguilar
// Creation Date: March 28th, 2024

package models

import (
	"strconv"
	"fmt"
)

type Row struct {
    name string
	email string
    age int
}

func NewRowFromSlice(input []string) *Row {
	intVar, _ := strconv.Atoi(input[2])
	return &Row{input[0], input[1], intVar}
}

func (row *Row) Print() {
	fmt.Println("[*] Row:")
	fmt.Println("\t" + row.name)
	fmt.Println("\t" + row.email)
}
