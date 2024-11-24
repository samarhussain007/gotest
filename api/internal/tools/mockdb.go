package tools

import (
	"time"
)

// mockDB struct simulates a mock database for testing purposes.
// It implements methods for retrieving user login details and coin balances.
type mockDB struct{}

// mockLoginDetails is a map that simulates the storage of login details in a database.
// The keys are usernames, and the values are the corresponding LoginDetails structs.
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123",  // Simulated auth token for user alex.
		Username:  "alex", // Simulated username for user alex.
	},
	"jason": {
		AuthToken: "1234", // Simulated auth token for user jason.
		Username:  "json", // Simulated username for user jason.
	},
}

// mockCoinDetails is a map that simulates the storage of coin balances in a database.
// The keys are usernames, and the values are the corresponding CoinDetails structs.
var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    123000, // Simulated coin balance for user alex.
		Username: "alex", // Simulated username for user alex.
	},
	"jason": {
		Coins:    20000,  // Simulated coin balance for user jason.
		Username: "json", // Simulated username for user jason.
	},
}

// GetUserLoginDetails is a method on the mockDB type that simulates retrieving user login details
// from a database. It takes a username string as input and returns a pointer to the corresponding
// LoginDetails if found, or nil if the user does not exist in the mock database.
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a delay in fetching data from the database
	time.Sleep(time.Second * 1)

	// Attempt to retrieve the LoginDetails from the mockLoginDetails map
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	// If the username is not found in the mock database, return nil
	if !ok {
		return nil
	}

	// Return the pointer to the retrieved LoginDetails
	return &clientData
}

// GetUserCoins is a method on the mockDB type that simulates retrieving user coin balance details
// from a database. It takes a username string as input and returns a pointer to the corresponding
// CoinDetails if found, or nil if the user does not exist in the mock database.
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate a delay in fetching data from the database
	time.Sleep(time.Second * 2)

	// Attempt to retrieve the CoinDetails from the mockCoinDetails map
	var coinData = CoinDetails{}
	coinData, ok := mockCoinDetails[username]

	// If the username is not found in the mock database, return nil
	if !ok {
		return nil
	}

	// Return the pointer to the retrieved CoinDetails
	return &coinData
}

// SetupDatabase is a method on the mockDB type that simulates the database setup process.
// In this case, it does not perform any real database setup, but returns nil to indicate success.
func (d *mockDB) SetupDatabase() error {
	// No real database setup required for mockDB
	return nil
}
