package main

import (
	"fmt"
	"time"
)

func generate(data string) <-chan string {
	chn := make(chan string)
	go func() {
		for {
			chn <- data
			time.Sleep(time.Duration(100 * time.Millisecond))
		}
	}()
	return chn
}

func fanin() {
	c1 := generate("Hello")
	c2 := generate("There")
	fanin := make(chan string)
	go func() {
		for {
			select {
			case str1 := <-c1:
				fanin <- str1
			case str2 := <-c2:
				fanin <- str2
			case <-time.After(100 * time.Millisecond):
				fanin <- "Timeout!"
				//return
				//default:
				//	fmt.Println("Default case")
			}
		}
	}()
	go func() {
		for {
			fmt.Println(<-fanin)
		}
	}()
}

func main() {
	fanin()
	time.Sleep(time.Duration(3 * time.Second))
}
