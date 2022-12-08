package main

import (
	"fmt"
	"log"
	"os"

	"github.com/infamousjoeg/authn-jwt-gitlab/internal/conjurapi"
)

// Environment variables to define:
// CONJUR_APPLIANCE_URL, CONJUR_ACCOUNT, CONJUR_AUTHN_JWT_SERVICE_ID,
// CONJUR_AUTHN_JWT_TOKEN, CONJUR_SECRET_ID

func checkEnvironmentVariables() error {
	// Check for environment variables and error if one is missing.
	variables := []string{
		"CONJUR_APPLIANCE_URL",
		"CONJUR_ACCOUNT",
		"CONJUR_AUTHN_JWT_SERVICE_ID",
		"CONJUR_AUTHN_JWT_TOKEN",
		"CONJUR_SECRET_ID",
	}

	for _, variable := range variables {
		if os.Getenv(variable) == "" {
			return fmt.Errorf("Environment variable %s is not set", variable)
		}
	}

	return nil
}

func main() {

	// Check for environment variables and error if one is missing.
	err := checkEnvironmentVariables()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Defining secret ID to retrieve, as per 12 factor
	// this is being accomplished via env variables.
	variableIdentifier := os.Getenv("CONJUR_SECRET_ID")

	// Loading configuration via defined Env vars:
	// CONJUR_APPLIANCE & CONJUR_ACCOUNT
	config, err := conjurapi.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load configuration. %s", err)
	}

	// Create a new Conjur client using environment variables
	conjur, err := conjurapi.NewClientFromEnvironment(config)
	if err != nil {
		log.Fatalf("Cannot create new client from environment variables. %s", err)
	}

	// Retrieve the secret value from Conjur
	secretValue, err := conjur.RetrieveSecret(variableIdentifier)
	if err != nil {
		log.Fatalf("Cannot retrieve secret value for %s. %s", variableIdentifier, err)
	}

	// Print the secret value to stdout
	fmt.Printf("%s", string(secretValue))
}
