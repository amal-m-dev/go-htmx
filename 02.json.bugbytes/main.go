package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	UserId    int    `json:"-"`
	ID        int    `json:"-"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	todoItem := &Todo{1, 1, "", false}

	todo, err := json.MarshalIndent(todoItem, "", "\t")

	if err != nil {
		log.Fatal("Marshal error:", err)
	}

	fmt.Printf(string(todo))
}
