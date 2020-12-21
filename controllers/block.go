package controllers

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
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
// @router /:blockHash [get]
func (c *BlockController) Get() {
	filename := "mock/block.json"
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
// @Description get all blocks
// @Success 200 {object} models.Block
// @router / [get]
func (c *BlockController) GetAll() {
	filename := "mock/blocks.json"
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
