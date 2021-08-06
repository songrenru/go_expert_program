package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	go HandleRequest(ctx)

	time.Sleep(time.Second*10)
}

func HandleRequest(ctx context.Context) {
	go WriteDatabase(ctx)
	go WriteRedis(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest Running.")
			time.Sleep(time.Second*2)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis Running.")
			time.Sleep(time.Second*2)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase Running.")
			time.Sleep(time.Second*2)
		}
	}
}
