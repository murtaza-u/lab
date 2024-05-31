package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Power struct {
	Type   string  `yaml:"type"`
	Impact float32 `yaml:"force"`
}

type Human struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`

	// without `inline`:
	// power:
	//	type: Super Human
	//	impact: 8.9
	//
	// with `inline`:
	// type: Super Human
	// impact: 8.9
	Power *Power `yaml:"power"`

	// without `flow`:
	// tags:
	//	- Greek
	//	- Nerd
	//
	// with `flow`:
	// tags: [Greek, Nerd]
	Tags []string `yaml:"tags,flow"`
}

func main() {
	h := &Human{
		Name: "Murtaza Udaipurwala",
		Age:  20,
		Power: &Power{
			Type:   "Super Human",
			Impact: 8.9,
		},
		Tags: []string{"Geek", "Nerd"},
	}
	out, err := yaml.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))

	// newH := new(Human)
	// err = yaml.Unmarshal(out, newH)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf(
	// 	"Name: %s | Age: %d | Tags: %v\n",
	// 	newH.Name, newH.Age, newH.Tags,
	// )
}
