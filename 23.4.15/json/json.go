package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int `json:"released"`
}

func main() {
	movies := []Movie{{
		Title: "love",
		Year:  3,
	},
	}
	data, err := json.Marshal(movies)

	if err != nil {
		log.Fatalf("json marshal failed, %s", err)
	}

	fmt.Println(string(data))
}
