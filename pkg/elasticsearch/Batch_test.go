package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"strconv"
	"testing"
	"time"
)

type Tweet struct {
	User     string                `json:"user"`
	Age      int                   `json:"age"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

//批量插入
func Batch(index string, type_ string, datas ...interface{}, t *testing.T) {

	bulkRequest := ElasticClient.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
func TestBatch(t *testing.T) {

	tweet1 := Tweet{User: "Jame1", Age: 23, Message: "Take One", Retweets: 1, Created: time.Now()}
	tweet2 := Tweet{User: "Jame2", Age: 32, Message: "Take Two", Retweets: 0, Created: time.Now()}
	tweet3 := Tweet{User: "Jame3", Age: 32, Message: "Take Three", Retweets: 0, Created: time.Now()}
	Batch("twitter", "doc", tweet1, tweet2, tweet3)
}
