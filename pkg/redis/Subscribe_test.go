package redis

import (
	"testing"
)

func TestSubscribe(t *testing.T) {
	NewRedis()
	channelTopic := "mychannel*"
	pubsub := RedisClient.PSubscribe(channelTopic)
	defer pubsub.Close()
	t.Log(pubsub.String())
	for i := 0; i < 100; i++ {
		subMsg, err := pubsub.ReceiveMessage()
		if err != nil {
			t.Error(err)
		}
		t.Log("Channel ", subMsg.Channel)
		t.Log("Payload ", subMsg.Payload)
	}
}
