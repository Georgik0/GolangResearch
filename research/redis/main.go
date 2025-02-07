package main

/*
	docker build -t redis-cluster .
	docker run -p 6379:6379 -p 6380:6380 -p 6381:6381 redis-cluster
*/

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Test building
func main() {
	ctx := context.Background()

	client := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{
				"localhost:6379",
				"localhost:6380",
				"localhost:6381",
			},
		},
	)

	status := client.Ping(ctx)
	if status.Err() != nil {
		panic(status.Err())
	}

	err := client.Set(ctx, "ключ", "значение", 0).Err()
	if err != nil {
		panic(err)
	}

	value := client.Get(ctx, "testKey")

	v, err := value.Result()

	fmt.Println("v =", v)
	fmt.Println("err =", err)
	fmt.Println("value.Err() =", value.Err())
}
