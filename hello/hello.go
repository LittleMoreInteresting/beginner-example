package main

import (
	"fmt"

	"beginner-example/hello/say"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(say.Hello("Tom"))
	fmt.Println(quote.Go())
}
