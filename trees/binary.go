package trees

import (
	"github.com/c2nc/snippets/errors"
)

var (
	errEmptyTree = errors.NewConst("btree is nil")
	errValueExist = errors.NewConst("value already exist in tree")
)

type btreeNode struct {
	val   int
	left  *btreeNode
	right *btreeNode
}

// NewBinaryTree - constructor
func NewBinaryTree() *btreeNode {
	return &btreeNode{}
}

// Insert - insert value into bt
func (tn *btreeNode) Insert(v int) error {
	switch {
	case tn == nil:
		return errEmptyTree
	case tn.val == v:
		return errValueExist
	case v > tn.val:
		if tn.right == nil {
			tn.right = &btreeNode{val: v}
			return nil
		}
		return tn.right.Insert(v)
	case v < tn.val:
		if tn.left == nil {
			tn.left = &btreeNode{val: v}
			return nil
		}
		return tn.left.Insert(v)
	default:
		tn.val = v
		return nil
	}
}

// Delete - remove value and node fom bt
func (tn *btreeNode) Delete(n int) {
	tn = tn.remove(n)
}

// Find - search value in slice
func (tn *btreeNode) Find(n int) *btreeNode {
	if tn == nil {
		return tn
	}

	switch {
	case n == tn.val:
		return tn
	case n < tn.val:
		return tn.Find(n)
	default:
		return tn.right.Find(n)
	}
}

// GetMin - get minimal value
func (tn *btreeNode) GetMin() int {
	if tn.left == nil {
		return tn.val
	}

	return tn.left.GetMin()
}

// GetMax - get maximum value
func (tn *btreeNode) GetMax() int {
	if tn.right == nil {
		return tn.val
	}

	return tn.right.GetMax()
}

// Iter - bt sorted iterator
func (tn *btreeNode) Iter() <-chan int {
	iter := make(chan int)

	go func() {
		tn.iterTreeSorted(iter)
		close(iter)
		return
	}()

	return iter
}

func (tn *btreeNode) iterTreeSorted(i chan int) {
	if tn != nil {
		tn.left.iterTreeSorted(i)
		i <- tn.val
		tn.right.iterTreeSorted(i)
	}
}

func (tn *btreeNode) inorderShift() *btreeNode {
	cur := tn
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (tn *btreeNode) remove(n int) *btreeNode {
	if tn == nil {
		return nil
	}

	switch {
	case n < tn.val:
		tn.left = tn.left.remove(n)
	case n > tn.val:
		tn.right = tn.right.remove(n)
	default:
		if tn.left == nil {
			return tn.right
		} else if tn.right == nil {
			return tn.left
		} else {
			t := tn.right.inorderShift()
			t.left = tn.left
			return tn.right
		}
	}
	return tn
}