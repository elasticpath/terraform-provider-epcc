package epcc

import (
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// cfg holds the current configuration.
var cfg Config

// initialisation for the epcc package
func init() {
	// Set default configuration values.
	cfg.BaseURL = "https://api.moltin.com/"
	cfg.ClientTimeout = 10 * time.Second
	cfg.RetryLimitTimeout = 30 * time.Second
	cfg.BetaFeatures = ""

	// If the package is being tested, ignore environment variables.
	if len(os.Args) > 1 && os.Args[1][:5] == "-test" {
		log.Println("Package initialised in testing mode.")
		log.Println("Environment variables will be ignored.")
		return
	}

	// Otherwise, process environment variables and store them in the global cfg.
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	if cfg.Credentials.ClientID == "" {
		log.Fatal("Required environment variable EPCC_CLIENT_ID not found")
	}
	if cfg.Credentials.ClientSecret == "" {
		log.Fatal("Required environment variable EPCC_CLIENT_SECRET not found")
	}
}

// Config is used to keep track of configuration in one place.
// fields tagged envconfig are read from environment variables.
// fields tagged default are default values.
type Config struct {
	Credentials struct {
		ClientID     string `envconfig:"EPCC_CLIENT_ID"`
		ClientSecret string `envconfig:"EPCC_CLIENT_SECRET"`
	}
	BaseURL           string `envconfig:"EPCC_API_BASE_URL"`
	BetaFeatures      string `envconfig:"EPCC_BETA_API_FEATURES"`
	ClientTimeout     time.Duration
	RetryLimitTimeout time.Duration
}
