package main

import (
	"github.com/theankitbhardwaj/latest-wayback-snapshot-redis/api"
	"github.com/theankitbhardwaj/latest-wayback-snapshot-redis/cache"
)

func main() {
	cache := cache.NewRedisClient("localhost:6379")

	server := api.NewAPIServer(":8080", *cache)

	server.Run()
}
