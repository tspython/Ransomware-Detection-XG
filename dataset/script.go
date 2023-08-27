package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"string"

	"github.com/pelletier/go-toml"
	vt "github.com/VirusTotal/vt-go"
)


func main() {

	// read IDS.txt
	file, err := os.open("IDS.txt")
	if(err != nil) {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) >= 2 {
			sh1 := parts[1]

			virusTotalResponse := callVirusTotalAPI(sha1)

			fmt.Println("SHA1: ", sha1)
			fmt.Println("VirusTotalResponse: ", virusTotalResponse)
		}
	}
}

func callVirusTotalAPI(sha1 string) string {
	apiKey := loadAPIKey("env.toml")

	client := vt.NewClient(*apiKey)


}

func loadAPIKey(filename string) (string, error) {
	config, err := toml.LoadFile(filename)
	if err != nil {
		return "", err
	}

	apiKey := config.Get("api.api_key").(string)
	return apiKey, nil
}

