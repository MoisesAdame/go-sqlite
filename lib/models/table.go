// Author: Moisés Adame-Aguilar
// Creation Date: March 28th, 2024

package models

import (
	"container/list"
)


type Table struct {
    names *list.List
	emails *list.List
    ages *list.List

    size int
}

func NewTable(size int) *Table {
	return &Table{list.New(), list.New(), list.New(), size}
}

func (table *Table) Size() int {
    return table.size
}
