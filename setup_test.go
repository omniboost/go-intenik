package intenik_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	asperion "github.com/omniboost/go-intenik"
)

var (
	client *asperion.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	username := os.Getenv("INTENIK_USERNAME")
	password := os.Getenv("INTENIK_PASSWORD")
	debug := os.Getenv("DEBUG")
	if err != nil {
		log.Fatal(err)
	}
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = asperion.NewClient(nil)
	client.SetUsername(username)
	client.SetPassword(password)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
