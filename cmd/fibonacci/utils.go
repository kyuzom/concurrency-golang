package main

import (
	"fmt"
	"strconv"
)

func fibonacci_v1(n int) []int {
	nums := []int{}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		b, a = a+b, b
		nums = append(nums, b)
	}
	return nums
}

func fibonacci_v2(a int, b int, n int) []int {
	nums := []int{}
	for i := 0; i < n; i++ {
		b, a = a+b, b
		nums = append(nums, b)
	}
	return nums
}

func fibonacci_v3(s int, n int) []int {
	nums := []int{}
	a, b := 0, 1
	for i := 0; i < (s + n); i++ {
		b, a = a+b, b
		if i >= s {
			nums = append(nums, b)
		}
	}
	return nums
}

type FibonacciData struct {
	nums []int
}

func fibonacci_v4(a int, b int, n int) <-chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			b, a = a+b, b
			ch <- b
		}
		close(ch)
	}()
	return ch
}

func fibonacci_v5(chi <-chan int) <-chan int {
	cho := make(chan int)
	go func() {
		for n := range chi {
			a, b := 0, 1
			for i := 0; i < n; i++ {
				b, a = a+b, b
				cho <- b
			}
			cho <- -1
		}
		close(cho)
	}()
	return cho
}

func usernum() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			var uis string
			fmt.Scanln(&uis)
			if uis == "q" {
				break
			}
			n, err := strconv.Atoi(uis)
			if err == nil {
				ch <- n
			} else {
				fmt.Println(err)
			}
		}
		close(ch)
	}()
	return ch
}
