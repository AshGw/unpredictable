package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func addUp(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := x + y
	fmt.Println("Add res:", result)
}

func subtract(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := x - y
	fmt.Println("Sub res:", result)
}

func unPredictable(nums *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(10 + 1)
		*nums = append(*nums, randNum)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var nums []int
	var wg sync.WaitGroup

	wg.Add(1)
	go unPredictable(&nums, &wg)

	for {
		if len(nums) >= 2 {
			wg.Add(2)
			go addUp(nums[len(nums)-1], nums[len(nums)-2], &wg)
			go subtract(nums[len(nums)-1], nums[len(nums)-2], &wg)
		}
		time.Sleep(1 * time.Second)
	}
}
