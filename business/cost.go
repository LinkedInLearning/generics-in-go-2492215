package business

// Cost computes the netto cost for a costumer for Solar and returns the value
func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

// Cost computes the netto cost for a costumer for Wind and returns the value
func (w Wind) Cost() float64 {
	return w.Netto * 0.3
}
