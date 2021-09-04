package main

import "log"

type Person struct {
	Name string
}

func main() {
	p := Person{
		Name: "Bob",
	}
	defer log.Println(p)

	p.Name = "Chris"
}
