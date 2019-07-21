package elasticsearch

import (
	"context"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	start := time.Now()
	info, code, err := ElasticClient.Ping(elasticHost[0]).Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	duration := time.Since(start)
	t.Log(duration)
	t.Log(code, info.Version.Number)
}
