package main

import (
	"fmt"
	"time"
)

func _002ForLoop() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("Sum from 1 to 5 is %d\n", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("Value of sum is %d\n", sum)
}

func _002Conditional() {
	name := "Jonas"
	if name == "Jonas" {
		fmt.Printf("My name is %s\n", name)
	}

	/* Shorthand notation */
	if sum := 2 + 2; sum == 4 {
		fmt.Printf("2 + 2 is indeed %d\n", sum)
	}

	if len(name) > 10 {
		fmt.Println("You have a really long name!")
	} else {
		fmt.Printf("Your name of %q is %d characters long\n", name, len(name))
	}

}

func _002Switch() {
	iceCreamFlavour := "strawberry"
	switch iceCreamFlavour {
	case "chocolate":
		fmt.Println("Mmm...chocolate")
	case "vanilla":
		fmt.Println("Nothing better than vanilla")
	case "strawberry":
		fmt.Println("Strawberry is the best berry")
	default:
		fmt.Println("You have some weird tastes")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func _002Defer() {
	/* Defer will execute the function after the surrounding function
	returns. Arguments to the deferred function are evaluated
	immediately, however. */
	defer fmt.Println("34")

	fmt.Println("12")
}
