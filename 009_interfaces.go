package main

import "fmt"

type speaker interface {
	speak() string
}

type quietSpeaker struct{}

func (q quietSpeaker) speak() string {
	return fmt.Sprint("Hello...")
}

type loudSpeaker struct{}

func (l loudSpeaker) speak() string {
	return fmt.Sprint("HELLO!")
}

func broadcast(speakers []speaker) {
	for _, s := range speakers {
		fmt.Println(s.speak())
	}
}

func _009Interfaces() {
	var q quietSpeaker
	var l loudSpeaker

	broadcast([]speaker{q, l})
}
