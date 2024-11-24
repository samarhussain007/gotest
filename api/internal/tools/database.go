package tools

import log "github.com/sirupsen/logrus"

// LoginDetails struct represents the login details of a user, including
// their authentication token and username.
type LoginDetails struct {
	AuthToken string // AuthToken stores the user's authentication token.
	Username  string // Username stores the user's name.
}

// CoinDetails struct represents the coin balance details of a user, including
// the total number of coins they have and their username.
type CoinDetails struct {
	Coins    int64  // Coins stores the number of coins the user has.
	Username string // Username stores the user's name.
}

// DatabaseInterface defines the methods that any type implementing this interface must have.
// It allows retrieving user login details, coin details, and setting up the database.
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails // Retrieves the login details of a user based on their username.
	GetUserCoins(username string) *CoinDetails         // Retrieves the coin balance of a user based on their username.
	SetupDatabase() error                              // Sets up the database connection or environment.
}

// NewDatabase is a function that initializes and sets up the database.
// It returns a pointer to a DatabaseInterface (which in this case is a mockDB), and an error if any.
func NewDatabase() (*DatabaseInterface, error) {
	// Initialize the mock database by creating an instance of mockDB and assigning it to the database variable.
	var database DatabaseInterface = &mockDB{}

	// Call SetupDatabase on the mock database to set up any necessary database configurations.
	var err error = database.SetupDatabase()

	// If there was an error setting up the database, log the error and return nil for the database pointer.
	if err != nil {
		log.Error(err)  // Log the error using logrus.
		return nil, err // Return nil for the database and the error encountered.
	}

	// Return a pointer to the database (in this case, mockDB) and a nil error if the setup was successful.
	return &database, err
}
