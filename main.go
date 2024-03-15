package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

	listCarsByBrand := &cobra.Command{
		Use:   "list",
		Short: "List the cars by brand",
		Run: func(cmd *cobra.Command, args []string) {
			if brand == "" {
				fmt.Println("The brand cannot be empty!")
				return
			}

			uppercasedBrand := strings.ToUpper(brand)

			filteredCars := getCarsByBrand(uppercasedBrand, cars)

			printCarsList(filteredCars)
		},
	}

	numberOfCarsByBrand.Flags().StringVarP(&brand, "brand", "b", "", "Get the number of cars by brand")
	listCarsByBrand.Flags().StringVarP(&brand, "brand", "b", "", "List the cars by brand")

	rootCmd.AddCommand(numberOfCarsByBrand)
	rootCmd.AddCommand(listCarsByBrand)
	rootCmd.Execute()
}

type Car struct {
	Brand      string `json:"brand"`
	Dealership string `json:"dealership"`
	Kilometers int    `json:"kilometers"`
	Price      int    `json:"price"`
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
		kilometers, _ := strconv.Atoi(record[2])
		price, _ := strconv.Atoi(record[3])
		data := Car{Brand: record[0], Dealership: record[1], Kilometers: kilometers, Price: price}

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
