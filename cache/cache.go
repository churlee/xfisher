package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	onlineCache *cache.Cache
)

func init() {
	onlineCache = cache.New(10*time.Minute, 30*time.Second)
}

func GetOnlineCache() *cache.Cache {
	return onlineCache
}
