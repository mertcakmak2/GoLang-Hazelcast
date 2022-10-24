package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hazelcast/hazelcast-go-client"
)

func main() {

	fmt.Println("hazelcast connecting...")
	ctx := context.TODO()
	config := hazelcast.Config{}
	config.Cluster.Network.SetAddresses("member1:5701", "member2:5701", "member3:5701")
	client, err := hazelcast.StartNewClientWithConfig(ctx, config)

	if err != nil {
		log.Fatal(err)
	}

	people, err := client.GetMap(ctx, "people")
	if err != nil {
		log.Fatal(err)
	}
	personName := "Jane Doe"
	if err = people.Set(ctx, personName, 30); err != nil {
		log.Fatal(err)
	}
	age, err := people.Get(ctx, personName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old.\n", personName, age)
	client.Shutdown(ctx)

	time.Sleep(time.Hour)
}
