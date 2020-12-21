package utils

import (
	"fmt"
	"github.com/json-iterator/go"
)

type ErrorMessage struct {
	Code    int
	Message string
}

var ERROR_MESSAGE = map[string]*ErrorMessage{
	"NO_BLOCK_HASH":       {400, "No block hash"},
	"NO_TRANSACTION_HASH": {400, "No transaction hash"},
}

func LogJson(result interface{}) {
	json, err := jsoniter.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("%s\n", json)
}
