package elasticsearch

import (
	"github.com/olivere/elastic"
	"time"
)

var elasticHost = []string{
	"http://192.168.137.128:9200/",
}

var ElasticClient *elastic.Client

//初始化
func init() {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetURL(elasticHost...),
		elastic.SetHealthcheckTimeoutStartup(5*time.Second))
	if err != nil {
		panic(err)
	}
}
