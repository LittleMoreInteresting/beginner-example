package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {
	/*ch := make(chan struct{}, 1)
	//ch := make(chan struct{})
	ch <- struct{}{}
	go func() {
		<-ch
		fmt.Println("receive message")
	}()

	ch1 := make(chan string, 1)
	go write(ch1)
	go read(ch1)*/
	/*in := make(chan int, 10)
	out1 := make(chan int, 10)
	out2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	go func() {
		for i := range out1 {
			fmt.Println("1:", i)
		}
	}()
	go func() {
		for i := range out2 {
			fmt.Println("2:", i)
		}
	}()
	fanOut(in, []chan int{out1, out2})

	time.Sleep(15 * time.Second)*/

	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)
	go func() {
		for {
			ch1 <- 1
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			ch1 <- 2
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
		time.Sleep(time.Second)
	}
}

func read(ch <-chan string) {
	fmt.Println("read from channel: " + <-ch)
}
func write(ch chan<- string) {
	ch <- "message"
	fmt.Println("write to channel ")
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanOut(ch <-chan int, out []chan int) {
	var wg sync.WaitGroup
	for v := range ch { // 从输入chan中读取数据
		for i := 0; i < len(out); i++ {
			wg.Add(1)
			go func(v int, idx int) {
				defer wg.Done()
				out[idx] <- v
			}(v, i)
		}
	}
	wg.Wait()
	for i := 0; i < len(out); i++ {
		close(out[i])
	}

}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}
