package test

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	_ "starcoin-api/routers"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

var esUrl = os.Getenv("STARCOIN_ES_URL")
var esUser = os.Getenv("STARCOIN_ES_USER")
var esPwd = os.Getenv("STARCOIN_ES_PWD")

// Test_Elasticsearch is a sample to run an endpoint test
func Test_Elasticsearch(t *testing.T) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			esUrl,
		},
		Username: esUser,
		Password: esPwd,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
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
