package environment

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

// DATABASEHOST ...
var (
	DATABASEHOST     string
	DATABASENAME     string
	DATABASEPORT     int
	DATABASEPASSWORD string
	DATABASETIMEOUT  int
	DATABASEUSERNAME string
	ENVIRONMENT      string
	APIPORT          string
)

func init() {
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/bickyeric/warda/.env")

	DATABASEHOST = databaseHost()
	DATABASENAME = databaseName()
	DATABASEPORT, _ = strconv.Atoi(databasePort())
	DATABASEPASSWORD = databasePassword()
	DATABASETIMEOUT = databaseTimeOut()
	DATABASEUSERNAME = databaseUsername()
	ENVIRONMENT = environment()
	APIPORT = apiPort()
}

// Load ...
func Load() {
	log.Println(fmt.Sprintf("starting %s mode", ENVIRONMENT))
}

// IsDevelopment ...
func IsDevelopment() bool {
	return ENVIRONMENT == "development"
}

// IsProduction ...
func IsProduction() bool {
	return ENVIRONMENT == "production"
}

func apiPort() string {
	return env("API_PORT")
}

func environment() string {
	return env("ENV")
}

func databaseHost() string {
	return env("DATABASE_HOST")
}

func databaseName() string {
	return env("DATABASE_NAME")
}

func databasePort() string {
	return env("DATABASE_PORT")
}

func databaseTimeOut() int {
	obj, _ := strconv.Atoi(env("DATABASE_TIMEOUT"))
	return obj
}

func databaseUsername() string {
	return env("DATABASE_USERNAME")
}

func databasePassword() string {
	return env("DATABASE_PASSWORD")
}

func env(key string) string {
	return os.Getenv(key)
}
