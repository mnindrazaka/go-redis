package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type User struct {
	Name  string
	Email string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connection to redis is success")

	jsonData, err := json.Marshal(User{Name: "M. Nindra Zaka", Email: "mnindrazaka@gmail.com"})

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := client.Set("user", jsonData, 0).Err(); err != nil {
		fmt.Println(err)
		return
	}

	val, err := client.Get("user").Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	var user User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Email)
}
