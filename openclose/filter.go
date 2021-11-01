package main

type FilterCriteria interface {
	Applies(product *Product) bool
}

type ColorCriteria struct {
	Color Color
}

func (cc ColorCriteria) Applies(product *Product) bool {
	return product.Color == cc.Color
}

type SizeCriteria struct {
	Size Size
}

func (tc SizeCriteria) Applies(product *SizeCriteria) bool {
	return product.Size == tc.Size
}

type ProductFilter struct{}

func (filter ProductFilter) FilterByCriteria(products []*Product, filterCriteria FilterCriteria) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if filterCriteria.Applies(product) {
			results = append(results, product)
		}
	}
	return results
}















type Conjunction struct {
	A FilterCriteria
	B FilterCriteria
}

func (conj Conjunction) Applies(product *Product) bool {
	return conj.A.Applies(product) && conj.B.Applies(product)
}
