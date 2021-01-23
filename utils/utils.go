package utils

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type ErrorMessage struct {
	Code    int
	Message string
}

var ERROR_MESSAGE = map[string]*ErrorMessage{
	"NO_BLOCK_HASH":       {400, "No block hash"},
	"INVALID_PAGE":        {400, "Invalid page info"},
	"NO_TRANSACTION_HASH": {400, "No transaction hash"},
	"NO_ADDRESS_HASH": {400, "No address hash"},
}

func LogJson(result interface{}) {
	json, err := jsoniter.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("%s\n", json)
}
