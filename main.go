package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dealership-cli"}

	cars, err := readCSV("ProcimoChallenge_Backend_Dealership.csv")

	if err != nil {
		log.Fatalln("arst", err)
	}

	var brand string

	numberOfCarsByBrand := &cobra.Command{
		Use:   "filter",
		Short: "Get the number of cars by the brand",
		Run: func(cmd *cobra.Command, args []string) {
			if brand == "" {
				fmt.Println("The brand cannot be empty!")
				return
			}

			uppercasedBrand := strings.ToUpper(brand)

			filteredCars := getCarsByBrand(uppercasedBrand, cars)

			fmt.Printf("Number of %s cars: %d", uppercasedBrand, len(filteredCars))
		},
	}

	numberOfCarsByBrand.Flags().StringVarP(&brand, "brand", "b", "", "Get the number of cars by brand")

	rootCmd.AddCommand(numberOfCarsByBrand)
	rootCmd.Execute()
}

type Car struct {
	Brand      string `json:"brand"`
	Dealership string `json:"dealership"`
	Kilometers string `json:"kilometers"`
	Price      string `json:"price"`
}

func readCSV(filename string) ([]Car, error) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		log.Fatalln("arst", err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return nil, err
	}

	var jsonData []Car

	for _, record := range records {
		data := Car{Brand: record[0], Dealership: record[1], Kilometers: record[2], Price: record[3]}

		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}

func getCarsByBrand(brand string, cars []Car) []Car {
	var filteredCars []Car

	for _, car := range cars {
		if car.Brand == brand {
			filteredCars = append(filteredCars, car)
		}
	}

	return filteredCars
}

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
