package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/redis.v4"
)

var (
	mongoHost = flag.String("mhost", "mongo_1", "")
	redisHost = flag.String("rhost", "redis_2", "")
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	flag.Parse()

	fmt.Printf("mongo:%s redis:%s\n", *mongoHost, *redisHost)

	r(*redisHost)
	m(*mongoHost)
}

func m(mhost string) {
	session, err := mgo.Dial(mhost)
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(
		&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"},
	)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

func r(rhost string) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", rhost),
		DB:   0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
