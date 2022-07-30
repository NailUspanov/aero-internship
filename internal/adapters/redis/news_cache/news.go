package news_cache

import (
	"aero-internship/gen/api"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

type NewsCache struct {
	client *redis.Client
}

func NewNewsCache(client *redis.Client) *NewsCache {
	return &NewsCache{client: client}
}

func (cache *NewsCache) Set(key string, post *api.NewsObject, exp time.Duration) {

	// serialize Post object to JSON
	jsonObject, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	cache.client.Set(key, jsonObject, exp*time.Second)
}

func (cache *NewsCache) Get(key string) *api.NewsObject {

	val, err := cache.client.Get(key).Result()
	if err != nil {
		return nil
	}

	news := api.NewsObject{}
	err = json.Unmarshal([]byte(val), &news)
	if err != nil {
		panic(err)
	}

	return &news
}
