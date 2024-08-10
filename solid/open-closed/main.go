/*
Open for extension, closed for modification
*/
package main

import "fmt"

type Color int

const (
	green Color = iota
	red
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f Filter) FilterByColor(products []Product, color Color) []*Product {
	res := make([]*Product, 0)

	for i, p := range products {
		if p.color == color {
			res = append(res, &products[i])
		}
	}

	return res
}

func (f Filter) FilterBySize(products []Product, size Size) []*Product {
	res := make([]*Product, 0)

	for i, p := range products {
		if p.size == size {
			res = append(res, &products[i])
		}
	}

	return res
}

//

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return spec.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return spec.size == p.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type OrSpecification struct {
	first, second Specification
}

func (spec OrSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) || spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	res := make([]*Product, 0)

	for i, p := range products {
		if spec.IsSatisfied(&p) {
			res = append(res, &products[i])
		}
	}

	return res
}

func main() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: medium}
	bag := Product{name: "Bag", color: red, size: medium}
	house := Product{name: "House", color: blue, size: large}

	products := []Product{apple, tree, bag, house}

	var f Filter

	fmt.Printf("Green products (old):\n")
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s\n", v.name)
	}

	fmt.Printf("Small products (old):\n")
	for _, v := range f.FilterBySize(products, small) {
		fmt.Printf(" - %s\n", v.name)
	}

	var bf BetterFilter
	gmSpec := AndSpecification{
		first:  ColorSpecification{green},
		second: SizeSpecification{medium},
	}
	fmt.Printf("Green and medium products (new):\n")
	for _, v := range bf.Filter(products, gmSpec) {
		fmt.Printf("- %s\n", v.name)
	}

	redOrSmallSpec := OrSpecification{
		first:  ColorSpecification{red},
		second: SizeSpecification{small},
	}
	fmt.Printf("Red or small products (new):\n")
	for _, v := range bf.Filter(products, redOrSmallSpec) {
		fmt.Printf("- %s\n", v.name)
	}
}
