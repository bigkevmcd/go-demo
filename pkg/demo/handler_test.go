package demo

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestHandleRequest(t *testing.T) {
	if os.Getenv("CI_PROJECT_DIR") != "" {
		t.Skip("no access to Redis")
	}
	testKey := "demo:test-key"
	testVal := randomString()
	r := createRedis(t, testKey)
	r.Set(context.TODO(), testKey, testVal, 0)
	h := New(Config{Key: testKey, Redis: r})
	srv := httptest.NewServer(h)
	defer srv.Close()

	res, err := srv.Client().Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	want := fmt.Sprintf("serving %s\n", testVal)
	if string(body) != want {
		t.Fatalf("response got %s, want %s", body, want)
	}
}

func createRedis(t *testing.T, key string) *redis.Client {
	opts, err := redis.ParseURL(redisURL())
	if err != nil {
		t.Fatal(err)
	}
	rdb := redis.NewClient(opts)

	t.Cleanup(func() {
		err := rdb.Del(context.TODO(), key).Err()
		if err != nil {
			t.Errorf("failed to delete key: %s", err)
		}

	})

	return rdb
}

func redisURL() string {
	if u := os.Getenv("TEST_REDIS_URL"); u != "" {
		return u
	}
	return "redis://localhost:6379/9"
}

func randomString() string {
	charset := "abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return fmt.Sprintf("%s", b)
}
