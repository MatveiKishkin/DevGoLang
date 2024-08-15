package main

import "fmt"

func Loop() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	/*************************************/

	array := []string{"gym", "basketball"}

	for key, value := range array {
		fmt.Println(key, value)
	}

	/*************************************/

	map_list := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range map_list {
		fmt.Println(key, value)
	}

	for _, value := range map_list {
		fmt.Println(value)
	}

	/*************************************/

	channel := make(chan int)

	fmt.Println(channel)

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()

	for item := range channel {
		fmt.Println(item)
	}
}
