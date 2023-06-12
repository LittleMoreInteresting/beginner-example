package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			fmt.Println("this goroutine exit !!!")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("do something")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
