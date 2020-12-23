package main

import (
	"fmt"

	"github.com/c2nc/snippets/sorting"
	"github.com/c2nc/snippets/trees"
)

func main() {
	t := trees.NewBinaryTree()
	for _, v := range sorting.GenIntSlice(20, 0) {
		t.Insert(v)
	}

	fmt.Println("tree before deleting")
	for n := range t.Iter() {
		fmt.Println(n)
	}

	t.Delete(10)

	fmt.Println("tree after deleting")
	for n := range t.Iter() {
		fmt.Println(n)
	}
}
