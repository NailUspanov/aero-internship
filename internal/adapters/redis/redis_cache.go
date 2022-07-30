package redis

import (
	"aero-internship/gen/api"
	"aero-internship/internal/adapters/redis/news_cache"
	"aero-internship/internal/adapters/redis/tags_cache"
	"github.com/go-redis/redis/v7"
	"time"
)

type NewsCache interface {
	Set(key string, post *api.NewsObject, exp time.Duration)
	Get(key string) *api.NewsObject
}

type TagsCache interface {
	Set(key string, post *api.Tag, exp time.Duration)
	Get(key string) *api.Tag
}

type CacheRedis struct {
	NewsCache
	TagsCache
}

func NewCacheRedis(client *redis.Client) *CacheRedis {
	return &CacheRedis{
		NewsCache: news_cache.NewNewsCache(client),
		TagsCache: tags_cache.NewTagsCache(client),
	}
}
