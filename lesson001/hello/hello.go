package main

import (
	"fmt"

	"beginner-example/lesson001/hello/say"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World（^-^）")
	fmt.Println(say.Hello("Tom"))
	fmt.Println(quote.Go())
}
