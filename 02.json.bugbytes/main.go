package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2/"

	respone, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer respone.Body.Close()

	if respone.StatusCode == http.StatusOK {
		todoItem := Todo{}

		decoder := json.NewDecoder(respone.Body)

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal("Decode error:", err)
		}

		//convert to json
		todo, err := json.Marshal(todoItem)

		if err != nil {
			log.Fatal("Marshal error:", err)
		}

		fmt.Printf(string(todo))
	}

}
