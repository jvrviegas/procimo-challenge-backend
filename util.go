package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func getUniqueBrands(cars []Car) []string {
	uniqueBrand := make(map[string]bool)
	var uniqueBrandsList []string

	for _, car := range cars {
		if _, found := uniqueBrand[car.Brand]; !found {
			uniqueBrand[car.Brand] = true
			uniqueBrandsList = append(uniqueBrandsList, car.Brand)
		}
	}

	return uniqueBrandsList
}

func listUniqueBrands(cars []Car) {
	uniqueBrands := getUniqueBrands(cars)

	fmt.Println("Unique brands: ")
	for _, uniqueBrand := range uniqueBrands {
		fmt.Println(uniqueBrand)
	}
}

func printCarsList(cars []Car) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Brand\t| Dealership \t| Kilometers \t| Price")
	fmt.Fprintln(w, "------\t-------------\t--------------\t-------")

	for _, c := range cars {
		fmt.Fprintf(w, "%s\t| %s\t| %d\t | %d\n", c.Brand, c.Dealership, c.Kilometers, c.Price)
	}

	w.Flush()
}

func formatCurrency(num int) string {
	numStr := fmt.Sprintf("%d", num)
	var formattedNum string

	for i, digit := range numStr {
		if i > 0 && (len(numStr)-i)%3 == 0 {
			formattedNum += ","
		}
		formattedNum += string(digit)
	}

	return "$" + formattedNum
}
