package main

import (
	"context"
	"fmt"

	"github.com/rajesh2412/go-chi-microservices/application"
)

func main() {

	app := application.New()

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "mykey", "Rajesh")
	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}
