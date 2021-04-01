package controllers

import (
	"html/template"
	"starcoin-explorer-api/db"
	"starcoin-explorer-api/utils"
	"strconv"
)

const TableBlock= "blocks"

// Operations about block
type BlockController struct {
	BaseController
}

// @Title Get
// @Description find block by blockHash
// @Param	network		path 	string	true		"the network you want to use"
// @Param	blockHash		path 	string	true		"the blockHash you want to get"
// @Success 200 {object} models.Block
// @Failure 403 :blockHash is empty
// @router /:network/hash/:blockHash [get]
func (c *BlockController) Get() {
	blockHash := template.HTMLEscapeString(c.GetString(":blockHash"))
	if blockHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_BLOCK_HASH"])
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"header.block_hash": blockHash,
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableBlock)

	c.Response(result, err)

}


// @Title Get By Height
// @Description find block by block height
// @Param	network		path 	string	true		"the network you want to use"
// @Param	blockHeight		path 	string	true		"the blockHeight you want to get"
// @Success 200 {object} models.Block
// @Failure 403 :blockHeight is empty
// @router /:network/height/:blockHeight [get]
func (c *BlockController) GetByHeight() {
	blockHeight := template.HTMLEscapeString(c.GetString(":blockHeight"))
	if blockHeight == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_BLOCK_HEIGHT"])
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"header.number": blockHeight,
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableBlock)

	c.Response(result, err)

}

// @Title GetAll
// @Description get all blocks
// @Param	network		path 	string	true		"the network you want to use"
// @Param	page		path 	int	true		"page number"
// @Success 200 {object} models.Block
// @router /:network/page/:page [get]
func (c *BlockController) GetAll() {
	page, _ := c.GetInt(":page")
	if !(page > 0) {
		c.Response(nil, nil, utils.ERROR_MESSAGE["INVALID_PAGE"])
		return
	}
	total, _ := strconv.Atoi(c.Ctx.Input.Query("total"))
	pageSize := 20
	from := (page - 1) * pageSize

	// elastic has limit of 10000, we need to get total first then use search_after
	if page > 500 && total == 0 {
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
			"from": 0,
			"size": 1,
			"sort": []map[string]interface{}{
				map[string]interface{}{
					"header.number": map[string]interface{}{
						"order": "desc",
					},
				},
			},
		}
		result, _ := db.Query(&query, esPrefix, TableBlock)
		var hits = result.(map[string]interface{})["hits"]
		var totalMap = hits.(map[string]interface{})["total"]
		var totalValue = totalMap.(map[string]interface{})["value"]
		total =  int(totalValue.(float64))
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": from,
		"size": pageSize,
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"header.number": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}
	if page > 500 && total > 0 {
		query["from"] = 0
		after := total - from
		query["search_after"] = []interface{}{ after }
	}
	result, err := db.Query(&query, esPrefix, TableBlock)

	c.Response(result, err)
}
