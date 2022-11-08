package business

import (
	"golang.org/x/exp/constraints"
)

// Cost computes the netto cost for a costumer for Solar and returns the value
func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

// Cost computes the netto cost for a costumer for Wind and returns the value
func (w Wind) Cost() float64 {
	return w.Netto * 0.3
}

// Number is either a floating point number or an integer
type Number interface {
	constraints.Float | constraints.Integer
}

// Cost multiplies usage with netto to compute the cost.
func Cost[T Number](usage, netto T) T {
	cost := usage * netto
	return cost
}
