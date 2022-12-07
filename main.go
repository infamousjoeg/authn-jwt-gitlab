package main

import (
	"fmt"
	"log"
	"os"

	"github.com/infamousjoeg/authn-jwt-gitlab/internal/conjurapi"
)

//ENV Vars for def in K8s:
// CONJUR_PASS_OBJECT, CONJUR_USER_OBJECT, CONJUR_AUTHN_TOKEN_FILE, CONJUR_APPLIANCE_URL
// CONJUR_VERSION, CONJUR_ACCOUNT

func main() {

	// Check for environment variables and error if one is missing.
	if os.Getenv("CONJUR_APPLIANCE_URL") == "" || os.Getenv("CONJUR_ACCOUNT") == "" {
		log.Fatalf("Both CONJUR_APPLIANCE_URL and CONJUR_ACCOUNT environment variables must be set.")
	}

	//Defining Username & Password objects to retrieve, as per 12 factor
	//this is being accomplished via env variables.
	variableIdentifier := os.Getenv("CONJUR_PASS_OBJECT")
	variableuserIdentifier := os.Getenv("CONJUR_USER_OBJECT")

	//Loading configuration via defined Env vars:
	//CONJUR_APPLIANCE, CONJUR_VERSION, CONJUR_ACCOUNT
	config, err := conjurapi.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load configuration. %s", err)
	}

	//Get Authorization token from shared store from sidecar
	conjur, err := conjurapi.NewClientFromEnvironment(config)
	if err != nil {
		log.Fatalf("Cannot create new client from environment variables. %s", err)
	}

	if os.Getenv("Conjur") == "True" {

		secretValue, err := conjur.RetrieveSecret(variableIdentifier)
		if err != nil {
			log.Fatalf("Could not retrieve secret. %s", err)
		}

		secretValueUser, err := conjur.RetrieveSecret(variableuserIdentifier)
		if err != nil {
			log.Fatalf("Could not retrieve secret. %s", err)
		}

		//Display Username & Password in log.
		log.Printf("%s:%s", "The Username Used: ", secretValueUser)
		log.Printf("%s:%s", "The Password Used: ", secretValue)

	} else {

		//Grab Password from Conjur
		secretValue, err := conjur.RetrieveSecret(variableIdentifier)
		if err != nil {
			log.Fatalf("Could not retrieve secret. %s", err)
		}

		//Grab Username from Conjur
		//secretValueUser, err := conjur.RetrieveSecret(variableuserIdentifier)
		//if err != nil {
		//	panic(err)
		//}

		//Display Username & Password in log.
		fmt.Printf("%s", secretValue)
	}

}
