package main

import (
	"fmt"

	"github.comLinkedInLearning/generics-in-go-2492215/business"
)

// main is our simple "playground" for the course.
// Note, that in production code, it is a good practice to keep the main function short.
func main() {
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar3k.Print())
	fmt.Println(solar2k.Print())
	fmt.Println(windwest.Print())
}
