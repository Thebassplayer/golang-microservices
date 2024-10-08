package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	go func(input chan string) {
		fmt.Println("Sending 1 to the channel")
		input <- "hello1"
		fmt.Println("Sending 2 to the channel")
		input <- "hello2"
		fmt.Println("Sending 3 to the channel")
		input <- "hello3"
		fmt.Println("Sending 4 to the channel")
		input <- "hello4"
	}(c)

	fmt.Println("Receiving from the channel")

	for gretting := range c {

		fmt.Println(gretting)
		fmt.Println("Gretting received")
	}

}
