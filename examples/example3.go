package main

import (
	"fmt"
	"github.com/shomali11/fridge"
	"time"
)

func main() {
	redisClient := fridge.NewRedisClient()
	client := fridge.NewClient(redisClient, fridge.WithDefaultDurations(time.Second, 2*time.Second))
	defer client.Close()

	fmt.Println(client.Put("food", "Pizza"))
	fmt.Println(client.Get("food"))

	time.Sleep(time.Second)

	fmt.Println(client.Get("food"))

	time.Sleep(2 * time.Second)

	fmt.Println(client.Get("food"))
	fmt.Println(client.Remove("food"))
}
