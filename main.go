package main

import (
	"flag"
	"fmt"
	"github.com/mikeAdamss/gtk/inventory"
	"github.com/mikeAdamss/gtk/models"
	"github.com/mikeAdamss/gtk/parsers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	command := flag.String("call", "", "the name of the rest endpoint we want to call.")
	flag.Parse()

	restInventory := inventory.GetInventory()

	if *command == "" {
		listCommandsThenExit(restInventory)
	}

	url := ""
	parser := ""
	headersObj := []models.Header{}

	for _, restObj := range restInventory.Functions {

		if restObj.Name == *command {
			url = restObj.Url
			parser = restObj.Parser
			headersObj = restObj.Headers
		}
	}

	if url == "" {
		log.Println("Aborting - no record of function: " + *command)
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	for _, headerObj := range headersObj {
		if headerObj.Env == false {
			req.Header.Set(headerObj.Key, headerObj.Value)
		} else {
			val := os.Getenv(headerObj.Key)
			req.Header.Set(headerObj.Key, val)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Sprintf("Request for func %s failed with status code %s", command, resp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch parser {
	case models.ArrayOfLines:
		parsers.ArrayOfLines(bodyBytes)
	default:
		fmt.Println(resp.StatusCode)
	}

}

func listCommandsThenExit(funcInventory models.Functions) {

	nameArray := []string{}
	funcArray := []string{}
	maxTextPad := 0

	for i := range funcInventory.Functions {

		if len(funcInventory.Functions[i].Name) > maxTextPad {
			maxTextPad = len(funcInventory.Functions[i].Name) + 1
		}

		nameArray = append(nameArray, funcInventory.Functions[i].Name)
		funcArray = append(funcArray, funcInventory.Functions[i].Description)
	}

	if maxTextPad < 8 {
		maxTextPad = 8
	}

	// TODO - nice display
	fmt.Printf("\nFUNCTION" + getPad(8, maxTextPad) + "DESCRIPTION\n")
	fmt.Printf("----------------------\n")

	// Display all the functions and descriptions
	for i := range funcInventory.Functions {
		f := funcInventory.Functions[i].Name
		d := funcInventory.Functions[i].Description
		pad := getPad(len(f), maxTextPad)
		fmt.Println(f + pad + d)
	}
	os.Exit(1)

}

func getPad(textLen, maxTextPad int) string {

	pad := ""
	for i := 0; i < (maxTextPad - textLen); i++ {
		pad = pad + " "
	}
	pad = pad + " | "

	return pad
}
