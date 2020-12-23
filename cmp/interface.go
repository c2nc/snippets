package cmp

// Lesser - comparable with `less than`
type Lesser interface {
	// Less - value less than
	Less(Lesser) bool
}

// Comparer - comparable with `less than` and `equal` values
type Comparer interface {
	Lesser
	// Equal - values are equal
	Equal(Comparer) bool
}
