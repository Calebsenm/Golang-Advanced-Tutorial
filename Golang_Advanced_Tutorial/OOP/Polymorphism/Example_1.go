package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}

	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax %d\n ", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax() // this where runtime polymorphism happend
	}

	return totalTax

}
