package main

import (
	"fmt"
	"time"
)

func main() {
	// collect the first 40 piece fo fibonacci numbers serially
	/*{
		start := time.Now()
		nums := fibonacci_v3(0, 40)
		finish := time.Since(start)
		fmt.Println(finish, nums)
	}*/
	// collect the first 40 piece fo fibonacci numbers concurrently
	/*{
		start := time.Now()
		ch := fibonacci_v4(0, 1, 40)
		nums := []int{}
		for n := range ch {
			nums = append(nums, n)
		}
		finish := time.Since(start)
		fmt.Println(finish, nums)
	}*/
	// infinite loop: get from user how many fibonacci numbers to collect (from the beginning)
	{
		start := time.Now()
		nums := []int{}
		fmt.Print("How many fibonacci numbers you would like to see? ")
		ch := fibonacci_v5(usernum())
		for n := range ch {
			if n != -1 {
				if nums == nil {
					start = time.Now()
					nums = []int{}
				}
				nums = append(nums, n)
			} else {
				finish := time.Since(start)
				fmt.Println(finish, nums)
				nums = nil
				fmt.Print("How many fibonacci numbers you would like to see? ")
			}
		}
	}
}
