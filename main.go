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

	var brand, dealership string
	var min, max int

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

	listCarsByMileageRange := &cobra.Command{
		Use:   "range",
		Short: "List the cars by mileage range",
		Run: func(cmd *cobra.Command, args []string) {
			if min <= -1 || max <= -1 {
				fmt.Println("The min and max cannot be empty!")
				return
			}

			filteredCars := getCarsByMileageRange(min, max, cars)

			printCarsList(filteredCars)
		},
	}

	getDealershipTotalAmount := &cobra.Command{
		Use:   "total",
		Short: "Get the total amount of a dealership",
		Run: func(cmd *cobra.Command, args []string) {
			if dealership == "" {
				fmt.Println("The dealership cannot be empty!")
				return
			}

			totalAmount := getTotalPriceByDealership(dealership, cars)

			fmt.Printf("Total amount of %s cars: %s\n", dealership, formatCurrency(totalAmount))
		},
	}

	numberOfCarsByBrand.Flags().StringVarP(&brand, "brand", "b", "", "Get the number of cars by brand")
	listCarsByBrand.Flags().StringVarP(&brand, "brand", "b", "", "List the cars by brand")
	listCarsByMileageRange.Flags().IntVar(&min, "min", -1, "Minimun mileage")
	listCarsByMileageRange.Flags().IntVar(&max, "max", -1, "Maximum mileage")
	getDealershipTotalAmount.Flags().StringVarP(&dealership, "dealership", "d", "", "Get the total amount of the cars given a dealership")

	rootCmd.AddCommand(numberOfCarsByBrand)
	rootCmd.AddCommand(listCarsByBrand)
	rootCmd.AddCommand(listCarsByMileageRange)
	rootCmd.AddCommand(getDealershipTotalAmount)
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

func getCarsByMileageRange(min, max int, cars []Car) []Car {
	var filteredCars []Car

	for _, car := range cars {
		if car.Kilometers <= max && car.Kilometers >= min {
			filteredCars = append(filteredCars, car)
		}
	}

	return filteredCars
}

func getTotalPriceByDealership(dealership string, cars []Car) int {
	total := 0

	for _, car := range cars {
		if dealership == car.Dealership {
			total += car.Price
		}
	}

	return total
}
