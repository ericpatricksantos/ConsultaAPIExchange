package Functions

import (
	"encoding/json"
	"fmt"
	"os"

	"main.go/Model"
)

func GetUrlAPI() Model.UrlExchange {
	file, _ := os.Open("settings.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	configuration := Model.UrlExchange{}

	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error: ", err)
	}

	return configuration
}
