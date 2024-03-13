package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	cars, err := readCSV("ProcimoChallenge_Backend_Dealership.csv")

	if err != nil {
		log.Fatalln("arst", err)
	}

	// Define flags
	brandPtr := flag.String("brand", "hello", "a string")

	// Parse flags
	flag.Parse()

	// Access flag values
	fmt.Println("brand:", *brandPtr)

	carsByBrand := getCarsByBrand(*brandPtr, cars)

	for _, car := range carsByBrand {
		fmt.Printf("%+v", car)
	}

	fmt.Printf("Number of %s cars: %d", *brandPtr, len(carsByBrand))
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
