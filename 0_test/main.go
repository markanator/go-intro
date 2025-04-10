package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	CanSpeak bool `json:"canSpeak"`
}

// Person extends Human by embedding it
type Person struct {
	Human        // Embedding Human struct
	Name  string `json:"name"`
}

func main() {
	p := Person{
		Human: Human{CanSpeak: true},
		Name:  "Alice",
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(data))
}
