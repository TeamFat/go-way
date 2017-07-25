package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Update this function to accept a channel parameter,
// and remove the return value.
func generateKey(channel chan int) {
	fmt.Println("Generating key")
	keys := []int{3, 5, 7, 11}
	key := keys[rand.Intn(len(keys))]
	time.Sleep(3 * time.Second)
	fmt.Println("Done generating")
	// Write the key to the channel instead of returning.
	channel <- key
}

func main() {
	rand.Seed(time.Now().Unix())
	// Create a channel.
	channel := make(chan int)
	// Create 3 more goroutines.
	for i := 0; i < 3; i++ {
		go generateKey(channel)
	}
	// Read and print keys from the channel.
	// This also causes the program to wait until 3
	// keys have been read.
	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println("All done!")
	//http://blog.teamtreehouse.com/goroutines-concurrency
}
