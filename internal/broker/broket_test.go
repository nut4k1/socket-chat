package broker

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestPublish(t *testing.T) {
	s := miniredis.RunT(t)
	Init(&FakeRedisConfig{Addr: s.Addr()})
	defer client.Close()

	Publish(
		"teststream",
		BrockerMessage{
			From:    "testuser1",
			To:      "testuser2",
			Message: "testmsg",
		},
	)

	r, _ := client.XInfoStream(context.Background(), "teststream").Result()
	if r.Length != 1 {
		t.Error("Stream is not containing message after publish")
	}
}

func testpublish() {
	Publish(
		"teststream",
		BrockerMessage{
			From:    "testuser1",
			To:      "testuser2",
			Message: "testmsg",
		},
	)
}
func testunsuregroup() {
	EnsureGroup("teststream", "testgroup", context.Background())
}

func TestConsume(t *testing.T) {
	s := miniredis.RunT(t)
	Init(&FakeRedisConfig{Addr: s.Addr()})
	defer client.Close()

	for range 11 {
		testpublish()
	}
	testunsuregroup()

	strms, _ := Consume("teststream", "testgroup", "testconsumer", context.Background())
	if len(strms[0].Messages) != 10 {
		t.Error("consumed msgs count != 10")
	}

	pending, _ := client.XPendingExt(context.Background(), &redis.XPendingExtArgs{
		Stream: "teststream",
		Group:  "testgroup",
		Count:  100,
		Idle:   0,
		Start:  "-",
		End:    "+",
	}).Result()

	if len(pending) != 10 {
		t.Error("PEL count != 10")
	}

	r, _ := client.XInfoStream(context.Background(), "teststream").Result()
	if r.Length != 11 {
		t.Error("stream msgs count != 11")
	}
}
