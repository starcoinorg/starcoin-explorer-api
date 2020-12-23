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

// Operations about block
type BlockController struct {
	BaseController
}

// @Title Get
// @Description find block by blockHash
// @Param	blockHash		path 	string	true		"the blockHash you want to get"
// @Success 200 {object} models.Block
// @Failure 403 :blockHash is empty
// @router /hash/:blockHash [get]
func (c *BlockController) Get() {
	blockHash := template.HTMLEscapeString(c.GetString(":blockHash"))
	fmt.Println("blockHash", blockHash)
	if blockHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_BLOCK_HASH"])
		return
	}
	var r map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"header.block_hash": blockHash,
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
		db.ES.Search.WithIndex("starcoin.blocks"),
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
// @Description get all blocks
// @Param	page		path 	int	true		"page number"
// @Success 200 {object} models.Block
// @router /:page [get]
func (c *BlockController) GetAll() {
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
		db.ES.Search.WithIndex("starcoin.blocks"),
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
