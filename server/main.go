package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Compote struct {
	Ingredient1 string `json:"ingredient1"`
	Ingredient2 string `json:"ingredient2"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Getting request...")
		tmp := Compote{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(data, &tmp)
		if !(tmp == Compote{}) {
			log.Printf("Mixing: %v and %v", tmp.Ingredient1, tmp.Ingredient2)
			fmt.Fprintf(w, "Prepared compote from %v and %v", tmp.Ingredient1, tmp.Ingredient2)
		}
	})

	fmt.Println("Start webserver on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
