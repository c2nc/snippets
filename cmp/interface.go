package cmp

// Comparable - comparable values interface
type Comparable interface {
	// LessThan - value less than
	LessThan(v interface{}) bool
	// EqualTo - values are equal
	EqualTo(v interface{}) bool
}