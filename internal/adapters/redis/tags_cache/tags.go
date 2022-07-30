package tags_cache

import (
	"aero-internship/gen/api"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

type TagsCache struct {
	client *redis.Client
}

func NewTagsCache(client *redis.Client) *TagsCache {
	return &TagsCache{client: client}
}

func (cache *TagsCache) Set(key string, tag *api.Tag, exp time.Duration) {

	// serialize object to JSON
	jsonObject, err := json.Marshal(tag)
	if err != nil {
		panic(err)
	}

	cache.client.Set(key, jsonObject, exp*time.Second)
}

func (cache *TagsCache) Get(key string) *api.Tag {

	val, err := cache.client.Get(key).Result()
	if err != nil {
		return nil
	}

	tag := api.Tag{}
	err = json.Unmarshal([]byte(val), &tag)
	if err != nil {
		panic(err)
	}

	return &tag
}
