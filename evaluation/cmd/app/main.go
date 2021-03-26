package main

import (
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (

	//retrieve node name
	node = os.Getenv("MY_NODE_NAME")
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "10.128.2.85:6379",
		Password: "",
		DB:       0,
	})

	// pong, err := client.Ping().Result()
	// log.Println(pong, err)

	err := set(client, node)
	if err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Wrote Node name " + node + " to redis")

}

func set(client *redis.Client, value string) error {
	err := client.Set("NodeName", value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
