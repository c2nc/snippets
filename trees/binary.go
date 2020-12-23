package trees

import (
	"github.com/c2nc/snippets/cmp"
	"github.com/c2nc/snippets/errors"
)

var (
	errEmptyTree = errors.NewConst("btree is nil")
	errValueExist = errors.NewConst("value already exist in tree")
)

type btreeNode struct {
	val   cmp.Comparable
	left  *btreeNode
	right *btreeNode
}

// NewBinaryTree - constructor
func NewBinaryTree() *btreeNode {
	return &btreeNode{}
}

// Insert - insert value into bt
func (tn *btreeNode) Insert(v cmp.Comparable) error {
	switch {
	case tn == nil:
		return errEmptyTree
	case tn.val == nil:
		tn.val = v
		return nil
	case tn.val.EqualTo(v):
		return errValueExist
	case tn.val.LessThan(v):
		if tn.right == nil {
			tn.right = &btreeNode{val: v}
			return nil
		}
		return tn.right.Insert(v)
	case v.LessThan(tn.val):
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
func (tn *btreeNode) Delete(v cmp.Comparable) {
	tn = tn.remove(v)
}

// Find - search value in slice
func (tn *btreeNode) Find(v cmp.Comparable) *btreeNode {
	if tn == nil {
		return tn
	}

	switch {
	case v.EqualTo(tn.val):
		return tn
	case v.LessThan(tn.val):
		return tn.Find(v)
	default:
		return tn.right.Find(v)
	}
}

// GetMin - get minimal value
func (tn *btreeNode) GetMin() cmp.Comparable {
	if tn.left == nil {
		return tn.val
	}

	return tn.left.GetMin()
}

// GetMax - get maximum value
func (tn *btreeNode) GetMax() cmp.Comparable {
	if tn.right == nil {
		return tn.val
	}

	return tn.right.GetMax()
}

// Iter - bt sorted iterator
func (tn *btreeNode) Iter() <-chan cmp.Comparable {
	iter := make(chan cmp.Comparable)

	go func() {
		tn.iterTreeSorted(iter)
		close(iter)
		return
	}()

	return iter
}

func (tn *btreeNode) iterTreeSorted(i chan cmp.Comparable) {
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

func (tn *btreeNode) remove(v cmp.Comparable) *btreeNode {
	if tn == nil {
		return nil
	}

	switch {
	case v.LessThan(tn.val):
		tn.left = tn.left.remove(v)
	case tn.val.LessThan(v):
		tn.right = tn.right.remove(v)
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