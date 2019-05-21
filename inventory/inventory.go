package inventory

import (
	"fmt"
	"github.com/mikeAdamss/gtk/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func GetInventory() models.Functions {
	iPath := os.Getenv("GTK_INVENTORY")
	if iPath == "" {
		log.Println("Aborting - no inventory was found at:" + iPath)
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
