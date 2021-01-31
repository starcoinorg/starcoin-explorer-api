package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

var ES *elasticsearch.Client
var esUrl = os.Getenv("STARCOIN_ES_URL")
var esUser = os.Getenv("STARCOIN_ES_USER")
var esPwd = os.Getenv("STARCOIN_ES_PWD")

func ConnectElasticSearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			esUrl,
		},
		Username: esUser,
		Password: esPwd,
	}
	var err error
	ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func Query(query *map[string]interface{}, esPrefix, table string) (interface{}, error){
	esIndex := fmt.Sprintf("%s.%s", esPrefix, table)

	// Build the request body.
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
		return nil, err
	}
	// Perform the search request.
	res, err := ES.Search(
		ES.Search.WithContext(context.Background()),
		ES.Search.WithIndex(esIndex),
		ES.Search.WithBody(&buf),
		ES.Search.WithTrackTotalHits(true),
		ES.Search.WithPretty(),
	)
	defer res.Body.Close()

	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return res, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s\n", err)
			return res, err
		} else {
			err = fmt.Errorf("%s: %s",
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			log.Println(err.Error())
			return res, err
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return r, err
}