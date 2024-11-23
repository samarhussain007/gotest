package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants defining the maximum prices for chicken and tofu
const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 3

func main() {
	// Create channels to communicate chicken and tofu deals
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	// List of websites to check for prices
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	// Launch goroutines to check prices for each website
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel) // Check chicken prices
		go checkTofuPrices(websites[i], tofuChannel)       // Check tofu prices
	}

	// Wait for and process messages from the channels
	sendMessage(chickenChannel, tofuChannel)
}

// Function to check chicken prices on a given website
// - website: Name of the website being checked
// - channel: Channel to send the website name if a deal is found
func checkChickenPrices(website string, channel chan string) {
	for {
		time.Sleep(time.Second * 1) // Simulate a delay in checking prices

		// Generate a random chicken price
		var chickenPrice = rand.Float32() * 20

		// If the price is within the acceptable range, send the website name to the channel
		if chickenPrice <= MAX_CHICKEN_PRICE {
			channel <- website
			break // Exit the loop once a deal is found
		}
	}
}

// Function to check tofu prices on a given website
// - website: Name of the website being checked
// - channel: Channel to send the website name if a deal is found
func checkTofuPrices(website string, channel chan string) {
	for {
		time.Sleep(time.Second * 1) // Simulate a delay in checking prices

		// Generate a random tofu price
		var tofuPrice = rand.Float32() * 20

		// If the price is within the acceptable range, send the website name to the channel
		if tofuPrice < MAX_TOFU_PRICE {
			channel <- website
			break // Exit the loop once a deal is found
		}
	}
}

// Function to handle messages from the chicken and tofu channels
// - chickenChannel: Channel sending chicken deal website names
// - tofuChannel: Channel sending tofu deal website names
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel: // Wait for a message from the chicken channel
		fmt.Println("Text Sent: Found deal on Chicken at,", website)
	case website := <-tofuChannel: // Wait for a message from the tofu channel
		fmt.Println("Text Sent: Found deal on tofu at,", website)
	}
}
