package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// contactInfo represents a struct to store basic contact information.
// - Name: The name of the contact.
// - Email: The email address of the contact.
type contactInfo struct {
	Name  string
	Email string
}

// purchaseInfo represents a struct to store purchase details.
// - Name: The name of the purchased item.
// - Price: The price of the item.
// - Amount: The quantity of the item purchased.
type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

// gasEngine represents the properties of a gas-powered engine.
// - gallons: The fuel tank capacity in gallons.
// - mpg: The miles per gallon the engine can achieve.
type gasEngine struct {
	gallons float32
	mpg     float32
}

// electricEngine represents the properties of an electric engine.
// - kwh: The battery capacity in kilowatt-hours.
// - mpkwh: The miles per kilowatt-hour the engine can achieve.
type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// car is a generic struct that can represent either a gas-powered or electric car.
// - Name: The name of the car.
// - brand: The brand of the car.
// - engine: The engine of the car, which can be either gasEngine or electricEngine.
type car[T gasEngine | electricEngine] struct {
	Name   string
	brand  string
	engine T
}

func main() {
	// Example slices to demonstrate sumSlice and isEmpty functions.
	var intSlice = []int32{1, 2, 3}
	var emptySlice = []int32{}

	// Calculate the sum of a slice of integers and print the result.
	fmt.Println(sumSlice(intSlice))

	// Check if slices are empty and print the results.
	fmt.Println(isEmpty(intSlice), isEmpty(emptySlice))

	// Load contactInfo and purchaseInfo data from JSON files (placeholder paths).
	// var contacts []contactInfo = loadJSON[contactInfo]("./file.path")
	// var purchases []purchaseInfo = loadJSON[purchaseInfo]("./file.path")

	// // Example of a car with a gas engine.
	// var gasCar = car[gasEngine]{
	// 	Name:  "Civic",
	// 	brand: "Honda",
	// 	engine: gasEngine{
	// 		mpg:     40,
	// 		gallons: 12.4,
	// 	},
	// }

	// // Example of a car with an electric engine.
	// var electricCar = car[electricEngine]{
	// 	Name:  "Civic",
	// 	brand: "Honda",
	// 	engine: electricEngine{
	// 		kwh:   40,
	// 		mpkwh: 12.4,
	// 	},
	// }

	// These examples demonstrate the use of generics for different data types.
}

// loadJSON is a generic function to load JSON data from a file into a slice of structs.
// - T: The type of data to load (either contactInfo or purchaseInfo).
// - filepath: The path to the JSON file.
// Returns a slice of data loaded from the JSON file.
func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	data, _ := os.ReadFile(filepath) // Read the file at the specified path.
	var loaded []T
	json.Unmarshal(data, &loaded) // Unmarshal the JSON data into the slice.
	return loaded
}

// sumSlice calculates the sum of a numeric slice.
// - T: The numeric type (int32, float32, float64).
// - slice: The slice of numbers to sum.
// Returns the sum of the slice.
func sumSlice[T int32 | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// isEmpty checks whether a slice is empty.
// - T: Any type of slice.
// - slice: The slice to check.
// Returns true if the slice is empty, otherwise false.
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
