package main

import (
	"context"
	"log"

	"github.com/kilichov-94/simple-redis-client/cache"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDb       int    = 0
)

func main() {
	ctx := context.Background()
	c := cache.New(RedisAddr, RedisPassword, RedisDb)
	if err := c.Ping(ctx); err != nil {
		log.Panic("failed to connect to Redis")
	}
	log.Print("connected to redis...")

	if err := c.Set(ctx, "username", "Nurmukhammad", 0); err != nil {
		log.Print("Err: Could not storage a value in Redis")
	}
	log.Print("Value in storage")

	val, err := c.Get(ctx, "username")
	if err != nil {
		log.Print("Not found!")
	}
	log.Printf("Result: %s\n", val)

}
