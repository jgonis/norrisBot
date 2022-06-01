package norrisFact

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const norrisURL = "https://api.chucknorris.io/jokes/random"

func GetNorrisFact(responseChannel chan<- string) {
	response, err := http.Get(norrisURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	byt, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(byt, &data)
	if err != nil {
		log.Fatal(err)
	}
	norrisFact := data["value"].(string)
	responseChannel <- norrisFact
}
