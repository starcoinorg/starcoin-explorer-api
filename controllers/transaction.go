package controllers

import (
	"html/template"
	"starcoin-explorer-api/db"
	"starcoin-explorer-api/utils"
)

const TableTransaction = "txn_infos"

// Operations about transaction
type TransactionController struct {
	BaseController
}

// @Title Get
// @Description find transaction by transactionHash
// @Param	network		path 	string	true		"the network you want to use"
// @Param	transactionHash		path 	string	true		"the transactionHash you want to get"
// @Success 200 {object} models.Transaction
// @Failure 403 :transactionHash is empty
// @router /:network/hash/:transactionHash [get]
func (c *TransactionController) Get() {
	transactionHash := template.HTMLEscapeString(c.GetString(":transactionHash"))
	if transactionHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_TRANSACTION_HASH"])
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"_id": transactionHash,
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableTransaction)

	c.Response(result, err)

}

// @Title GetAll
// @Description get all transactions
// @Param	network		path 	string	true		"the network you want to use"
// @Param	page		path 	int	true		"page number"
// @Success 200 {object} models.Transaction
// @router /:network/page/:page [get]
func (c *TransactionController) GetAll() {
	page, _ := c.GetInt(":page")
	if !(page > 0) {
		c.Response(nil, nil, utils.ERROR_MESSAGE["INVALID_PAGE"])
		return
	}
	pageSize := 20
	from := (page - 1) * pageSize


	query := map[string]interface{}{
		"query": map[string]interface{}{
			//"match_all": map[string]interface{}{},
			"range": map[string]interface{}{
				"transaction_index": map[string]interface{}{
					"gt": 0,
				},
			},
		},
		"from": from,
		"size": pageSize,
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"timestamp": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableTransaction)

	c.Response(result, err)
}

// @Title Get Transactions by Address
// @Description find transactions by address hash
// @Param	network		path 	string	true		"the network you want to use"
// @Param	addressHash		path 	string	true		"the addressHash you want to get"
// @Success 200 {object} models.Transaction
// @Failure 403 :addressHash is empty
// @router /:network/byAddress/:addressHash [get]
func (c *TransactionController) GetListByAddress() {
	addressHash := template.HTMLEscapeString(c.GetString(":addressHash"))
	if addressHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_ADDRESS_HASH"])
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"user_transaction.raw_txn.sender": addressHash,
			},
		},
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"user_transaction.raw_txn.sequence_number": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableTransaction)

	c.Response(result, err)

}

// @Title Get Transactions by Block
// @Description find transactions by block hash
// @Param	network		path 	string	true		"the network you want to use"
// @Param	blockHash		path 	string	true		"the blockHash you want to get"
// @Success 200 {object} models.Transaction
// @Failure 403 :blockHash is empty
// @router /:network/byBlock/:blockHash [get]
func (c *TransactionController) GetListByBlock() {
	blockHash := template.HTMLEscapeString(c.GetString(":blockHash"))
	if blockHash == "" {
		c.Response(nil, nil, utils.ERROR_MESSAGE["NO_ADDRESS_HASH"])
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"block_hash": blockHash,
			},
		},
		"sort": []map[string]interface{}{
			map[string]interface{}{
				"transaction_index": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	result, err := db.Query(&query, esPrefix, TableTransaction)

	c.Response(result, err)

}
