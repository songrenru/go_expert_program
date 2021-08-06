package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "parameter", 1)
	go HandleRequest(ctx)

	time.Sleep(time.Second*10)
}

func HandleRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(time.Second*2)
		}
	}
}
