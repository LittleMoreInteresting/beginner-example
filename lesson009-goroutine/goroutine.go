package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func echo_chan(tag string, stop chan struct{}) {
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
func echo_ctx(tag string, ctx context.Context) {
	begin := true
	go func() {
		select {
		case <-ctx.Done():
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
	go echo_chan("1", stop)
	go echo_chan("2", stop)
	time.Sleep(5 * time.Second)
	close(stop)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go echo_ctx("ctx1", ctx)
	go echo_ctx("ctx2", ctx)
	time.Sleep(6 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Go 1]: this is Goroutine 1")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Go 2]: this is Goroutine 1")
	}()
	wg.Wait()
	fmt.Println("end")
}
