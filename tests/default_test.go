package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"starcoin-explorer-api/db"
	_ "starcoin-explorer-api/routers"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	db.ConnectElasticSearch()
}

// TestElasticsearch is a sample to run an endpoint test
func TestElasticsearch(t *testing.T) {
	res, err := db.ES.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var objmap map[string]interface{}
	if err := json.Unmarshal(body, &objmap); err != nil {
		log.Fatal(err)
	}
	//utils.LogJson(objmap)
	assert.Equal(t, objmap["cluster_name"], "starcoin-elasticsearch", "es cluster_name should be: starcoin-elasticsearch")
}

// TestMock is a sample to get json from mock file
func TestMock(t *testing.T) {
	filename := "mock/transactions.json"
	var objmap map[string]interface{}
	var jsonBlob []byte
	var err error
	jsonBlob, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	err = jsoniter.Unmarshal(jsonBlob, &objmap)
	if err != nil {
		fmt.Println("Unmarshal json fail: ", err.Error())
	}
	//utils.LogJson(objmap)
	assert.Equal(t, objmap["id"], "2", "id of mock/transactions.json should be: 2")
}
