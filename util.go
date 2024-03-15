package main

import "fmt"

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

func printCar(car Car) {
	if len(car.Dealership) <= 10 {
		fmt.Printf("%s \t | %s \t\t | %d \t | %d | \n", car.Brand, car.Dealership, car.Kilometers, car.Price)
	} else {
		fmt.Printf("%s \t | %s \t | %d \t | %d | \n", car.Brand, car.Dealership, car.Kilometers, car.Price)
	}
}
