package main

import (
	"fmt"
	"time"
)

func echo(tag string, stop chan struct{}) {
	begin := true
	go func() {
		select {
		case <-stop:
			fmt.Println(tag + "stop!")
			begin = false
		}
	}()
	for begin {
		fmt.Println("[" + tag + "]" + time.Now().String())
		time.Sleep(time.Second)
	}
}

func main() {
	stop := make(chan struct{})
	go echo("1", stop)
	go echo("2", stop)
	time.Sleep(5 * time.Second)
	close(stop)
	fmt.Println("end")
}
