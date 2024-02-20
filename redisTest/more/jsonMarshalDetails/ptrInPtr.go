package jsonMarshalDetails

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type Agent struct {
	Proxy  string `json:"proxy,omitempty"`
	UrlPos int    `json:"url_pos,omitempty"`
}

type Player struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Agent *Agent `json:"agent,omitempty"`
}

func redisWork() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	p := &Player{
		Name: "xiaoming",
		Age:  2,
		Agent: &Agent{
			Proxy:  "google.com",
			UrlPos: 15,
		},
	}

	bytes, _ := json.Marshal(&p)
	c1.Do("Set", "player", bytes)
	fmt.Println("marshal bytes ", string(bytes))

	nBytes, _ := redis.Bytes(c1.Do("Get", "player"))
	fmt.Println("marshal nBytes ", string(nBytes))

	np := &Player{}
	json.Unmarshal(nBytes, np)
	fmt.Println(np)

}
