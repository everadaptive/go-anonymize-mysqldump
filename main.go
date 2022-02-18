package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/humanmade/go-anonymize-mysqldump/pkg"
	"github.com/sirupsen/logrus"
)

// Many thanks to https://stackoverflow.com/a/47515580/1454045
func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to info
	if !ok {
		lvl = "info"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.InfoLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}

func main() {
	config := parseArgs()

	lines := pkg.SetupAndProcessInput(config, os.Stdin)

	for line := range lines {
		fmt.Print(<-line)
	}
}

func parseArgs() pkg.Config {
	parser := argparse.NewParser("anonymize-mysqldump", "Reads SQL from STDIN and replaces content for anonymity based on the provided config.")
	configFilePath := parser.String("c", "config", &argparse.Options{Required: true, Help: "Path to config.json"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	return readConfigFile(*configFilePath)
}

func readConfigFile(filepath string) pkg.Config {
	jsonConfig, err := ioutil.ReadFile(filepath)
	if err != nil {
		logrus.Fatal(err)
	}

	var decoded pkg.Config
	jsonReader := strings.NewReader(string(jsonConfig))
	jsonParser := json.NewDecoder(jsonReader)
	jsonParser.Decode(&decoded)
	return decoded
}
