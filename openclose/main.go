package main

import (
	"fmt"
)

func main() {
	products := []*Product{
		{Name: "car", Color: green, Size: medium},
		{Name: "bike", Color: green, Size: small},
		{Name: "scooter", Color: red, Size: small},
	}

	filter := ProductFilter{}
	filteredProducts := filter.FilterByCriteria(products, ColorCriteria{Color: green})
	// filteredProducts := filter.FilterByCriteria(products, SizeCriteria{Size: small})

	fmt.Println("filtered products\n")
	for _, product := range filteredProducts {
		fmt.Println(product.Name)
	}
}
