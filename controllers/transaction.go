package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"starcoin-api/db"
	"starcoin-api/utils"
)

// Operations about transaction
type TransactionController struct {
	BaseController
}

// @Title Get
// @Description find transaction by transactionHash
// @Param	transactionHash		path 	string	true		"the transactionHash you want to get"
// @Success 200 {object} models.Transaction
// @Failure 403 :transactionHash is empty
// @router /hash/:transactionHash [get]
func (c *TransactionController) Get() {
	transactionHash := template.HTMLEscapeString(c.GetString(":transactionHash"))
	fmt.Println("transactionHash", transactionHash)
	if transactionHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_TRANSACTION_HASH"])
		return
	}
	var r map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"transaction_hash": transactionHash,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	log.Print(query)

	// Perform the search request.
	res, err := db.ES.Search(
		db.ES.Search.WithContext(context.Background()),
		db.ES.Search.WithIndex("starcoin.txn_infos"),
		db.ES.Search.WithBody(&buf),
		db.ES.Search.WithTrackTotalHits(true),
		db.ES.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	utils.LogJson(r)

	c.Response(r, err)

}

// @Title GetAll
// @Description get all transactions
// @Param	page		path 	int	true		"page number"
// @Success 200 {object} models.Transaction
// @router /:page [get]
func (c *TransactionController) GetAll() {
	page, _ := c.GetInt(":page")
	fmt.Printf("page=%d\n", page)
	if !(page > 0) {
		c.Response(nil, nil, utils.ERROR_MESSAGE["INVALID_PAGE"])
		return
	}
	pageSize := 10
	from := (page - 1) * pageSize
	fmt.Printf("from=%d size=%d\n", from, pageSize)
	var r map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": from,
		"size": pageSize,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	log.Print(query)

	// Perform the search request.
	res, err := db.ES.Search(
		db.ES.Search.WithContext(context.Background()),
		db.ES.Search.WithIndex("starcoin.txn_infos"),
		db.ES.Search.WithBody(&buf),
		db.ES.Search.WithTrackTotalHits(true),
		db.ES.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	utils.LogJson(r)

	c.Response(r, err)
}
