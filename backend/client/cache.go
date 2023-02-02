package client

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func NewCache() {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(60*time.Minute))
	if err != nil {
		fmt.Println("缓存实例化失败：", err)
	}
	Cache = cache
}
