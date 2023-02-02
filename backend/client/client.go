package client

import (
	"github.com/allegro/bigcache/v3"
)

var (
	Cache *bigcache.BigCache
)

func init() {
	NewCache()
}
