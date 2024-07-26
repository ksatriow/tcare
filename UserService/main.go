// main.go

package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		err := rdb.Set(ctx, "user", "UserService", 0).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		val, err := rdb.Get(ctx, "user").Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, val)
	})

	fmt.Println("UserService is running on :8081")
	http.ListenAndServe(":8081", nil)
}
