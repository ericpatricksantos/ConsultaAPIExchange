package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"main.go/Functions"
	"main.go/Model"
)

func main() {

	response, err := http.Get(Functions.GetUrlAPI().UrlAPI)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Model.OrderBook
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("Symbol: " + responseObject.Symbol)
	fmt.Println("Tamanho do Bids: " + strconv.Itoa(len(responseObject.Bids)))
	fmt.Println("Tamanho do Asks: " + strconv.Itoa(len(responseObject.Asks)))

}
