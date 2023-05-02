package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Message struct {
	Name string
	Age int
}

func main() {

m := Message{"Panais", 24}

var res Message

b, err := json.Marshal(m)

client := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

ctx := context.Background()


seterr := client.Set(ctx, "myobj", string(b), 0).Err()
if seterr != nil {
	panic(err)
}

val, geterr := client.Get(ctx, "myobj").Result()
if geterr != nil {
	panic(geterr)
}

unerr := json.Unmarshal([]byte(val), &res)
if unerr != nil {
	panic(unerr)
}


fmt.Println(res.Name)

fmt.Println("Done")
//fmt.Println("foo", val)
}
