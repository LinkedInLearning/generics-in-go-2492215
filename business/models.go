package business

import (
	"fmt"
	"math"

	"golang.org/x/exp/slices"
)

// Solar handles all the different energy offers powered by solar.
type Solar struct {
	Name  string
	Netto float64
}

// SolarSlice is a defined slice of Solar.
type SolarSlice []Solar

// Wind handles all the different energy offers powered by wind.
type Wind struct {
	Name  string
	Netto float64
}

// Energy is either of type Wind or Solar
type Energy interface {
	Solar | Wind
	Cost() float64
}

// Print prints the information for a solar product.
// The string is enriched with the required kineteco legal information.
func (s *Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *s)
}

// Print prints the information for a wind product.
// The string is enriched with the required kineteco legal information.
func (w *Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *w)
}

// PrintGeneric returns any type as string.
// The string is enriched with the required Kineteco legal information.
func PrintGeneric[T Energy](t T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, t)
}

// PrintSlice prints a slice of any type to the standard output.
// Each item is enriched with its position and the Kineteco specific string.
func PrintSlice[T Energy](tt []T) {
	fmt.Printf("type of tt: %T\n", tt)
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

// PrintSlice prints a slice of any type to the standard output.
// Each item is enriched with its position and the Kineteco specific string.
func PrintSlice2[T Energy, S ~[]T](tt S) {
	fmt.Printf("type of tt: %T\n", tt)
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

func SortByCost[T Energy](a []T) {
	slices.SortFunc(a, func(a, b T) bool {
		return a.Cost() < b.Cost() || math.IsNaN(a.Cost()) && !math.IsNaN(b.Cost())
	})
}

var kinetecoPrint string = "Kineteco Deal:"
