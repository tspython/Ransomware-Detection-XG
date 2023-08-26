package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
	vt "github.com/VirusTotal/vt-go"
)

var sha256 = flag.String("sha256", "", "SHA-256 of some file")

func main() {
	flag.Parse()

	if *sha256 == "" {
		fmt.Println("Must pass the --sha256 argument.")
		os.Exit(0)
	}

	// Load the API key from env.toml
	apiKey, err := loadAPIKey("env.toml")
	if err != nil {
		log.Fatal(err)
	}

	client := vt.NewClient(apiKey)

	file, err := client.GetObject(vt.URL("files/%s", *sha256))
	if err != nil {
		log.Fatal(err)
	}

	ls, err := file.GetTime("last_submission_date")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s was submitted for the last time on %v\n", file.ID(), ls)
}

func loadAPIKey(filename string) (string, error) {
	config, err := toml.LoadFile(filename)
	if err != nil {
		return "", err
	}

	apiKey := config.Get("api.api_key").(string)
	return apiKey, nil
}

