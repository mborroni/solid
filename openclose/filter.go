package main

type ProductFilter struct {}

func (filter ProductFilter) FilterByColor(products []*Product, color Color) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if product.Color == color {
			results = append(results, product)
		}
	}
	return results
}

func (filter ProductFilter) FilterBySize(products []*Product, size Size) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if product.Size == size {
			results = append(results, product)
		}
	}
	return results
}

func (filter ProductFilter) FilterByColorAndSize(products []*Product, color Color, size Size) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if  product.Color == color && product.Size == size {
			results = append(results, product)
		}
	}
	return results
}