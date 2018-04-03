package main

import (
	"context"
	"fmt"
	"time"
)

type key int

var key1 key = 0
var key2 key = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key1, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case a, ok := <-ctx.Done():
			fmt.Println(a, ok)
			//get value
			fmt.Println(ctx.Value(key1), "is cancel")
			valueCtx := context.WithValue(ctx, key2, "add1 value")
			go watch1(valueCtx)
			return
		default:
			//get value
			fmt.Println(ctx.Value(key1), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}
func watch1(ctx context.Context) {
	fmt.Println(ctx)
	fmt.Println(ctx.Value(key1))
	fmt.Println(ctx.Value(key2))
}

// add value int goroutine
// add value int goroutine
// add value int goroutine
// add value int goroutine
// add value int goroutine
// {} false
// add value is cancel
// context.Background.WithCancel.WithValue(0, "add value").WithValue(1, "add1 value")
// add value
// add1 value
