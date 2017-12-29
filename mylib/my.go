package mylib

// Adder is a struct that helps you add numbers.
type Adder struct {
	x int
}

// Op performs the addition operator on the adder value and y.
func (a Adder) Op(y int) int {
	return a.x + y
}

// NewAdder makes a new struct that adds numbers.
func NewAdder(x int) Adder {
	return Adder{x}
}

// Subtractor is a struct that helps you subtract numbers.
type Subtractor struct {
	x int
}

// Op performs subtraction of y from x.
func (a Subtractor) Op(y int) int {
	return a.x - y
}

// NewSubtractor makes a new struct that subtracts numbers.
func NewSubtractor(x int) Subtractor {
	return Subtractor{x}
}
