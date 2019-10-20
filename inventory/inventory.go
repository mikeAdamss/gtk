package inventory

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mikeAdamss/gtk/models"
	"gopkg.in/yaml.v2"
)

func GetInventory() models.Functions {
	iPath := os.Getenv("GTK_INVENTORY")
	if iPath == "" {
		log.Println("Aborting - no inventory was found. Have you set the environment variable GTK_INVENTORY? This should point to your inventory file." + iPath)
		os.Exit(1)
	}

	yamlFile, err := ioutil.ReadFile(iPath)
	if err != nil {
		panic(err)
	}

	funcInventory := models.Functions{}
	err = yaml.Unmarshal(yamlFile, &funcInventory)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// Fail if someone has forgot to add the url
	// TODO - far more validation
	for i := range funcInventory.Functions {
		f := funcInventory.Functions[i]
		if f.Url == "" {
			log.Fatal(fmt.Sprintf("Aborting. The function '%s' is missing a value for url.", f.Name), "")
		}
	}

	return funcInventory
}
