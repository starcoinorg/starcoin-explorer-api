package controllers

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
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
// @router /:transactionHash [get]
func (c *TransactionController) Get() {
	filename := "mock/transaction.json"
	var output map[string]interface{}
	var jsonBlob []byte
	var err error
	jsonBlob, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	err = jsoniter.Unmarshal(jsonBlob, &output)
	if err == nil {
		fmt.Println("mock file:")
		utils.LogJson(output)
	}
	c.Response(output, err)
}

// @Title GetAll
// @Description get all transactions
// @Success 200 {object} models.Transaction
// @router / [get]
func (c *TransactionController) GetAll() {
	filename := "mock/transactions.json"
	var output map[string]interface{}
	var jsonBlob []byte
	var err error
	jsonBlob, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	err = jsoniter.Unmarshal(jsonBlob, &output)
	if err == nil {
		fmt.Println("mock file:")
		utils.LogJson(output)
	}
	c.Response(output, err)
}
