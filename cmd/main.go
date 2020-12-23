package main

import (
	"fmt"

	"github.com/c2nc/snippets/sorting"
	"github.com/c2nc/snippets/trees"
)

type IntVal int

func (c IntVal) Less(v interface{}) bool {
	if cv, ok := v.(IntVal); ok {
		return c < cv
	}

	return false
}

func (c IntVal) Equal(v interface{}) bool {
	if cv, ok := v.(IntVal); ok {
		return c == cv
	}

	return false
}

func main() {
	t := trees.NewBinaryTree()
	for _, v := range sorting.GenIntSlice(20, 0) {
		t.Insert(IntVal(v))
	}

	fmt.Println("tree before deleting")
	for n := range t.Iter() {
		fmt.Println(n)
	}

	t.Delete(IntVal(10))

	fmt.Println("tree after deleting")
	for n := range t.Iter() {
		fmt.Println(n)
	}
}
