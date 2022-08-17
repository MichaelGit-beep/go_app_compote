package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Fruits struct {
	Ingredient1 string `json:"ingredient1"`
	Ingredient2 string `json:"ingredient2"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	log.Println(host, port)
	var socket string
	if host != "" && port != "" {
		socket = fmt.Sprintf("http://%s:%s", host, port)
	} else {
		socket = "http://localhost:8080"
	}
	log.Println("Connecting to: ", socket)
	fruits := []string{"orange", "apple", "pineapple", "banana",
		"cherry", "lime", "raspberry", "plum", "guava", "mango", "grapes"}

	for {
		time.Sleep(2 * time.Second)
		d, err := json.Marshal(Fruits{fruits[rand.Intn(len(fruits))], fruits[rand.Intn(len(fruits))]})
		if err != nil {
			log.Println(err)
		}
		req, err := http.Post(socket, "application/json", bytes.NewReader(d))
		if err != nil {
			log.Println(err)
			continue
		}
		responce, err := ioutil.ReadAll(req.Body)
		if string(responce) != "" {
			log.Println(string(responce))
		} else {
			log.Println("Getting empty responce")
		}
	}
}
