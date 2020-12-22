package db

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

var ES *elasticsearch.Client
var esUrl = os.Getenv("STARCOIN_ES_URL")
var esUser = os.Getenv("STARCOIN_ES_USER")
var esPwd = os.Getenv("STARCOIN_ES_PWD")

func ConnectElasticsearch() {
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
