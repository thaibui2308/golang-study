package controller

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"example.com/goinpractice/cli-app/urfave-cli/models"
)

func GetCountryCovidCase(country string) models.Response {
	response, err := http.Get("https://corona.lmao.ninja/v2/countries/" + country + "?yesterday&strict&query%20")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
