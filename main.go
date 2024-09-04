package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/rajesh2412/go-chi-microservices/application"
)

func main() {

	app := application.New()
	// with the support of signal notifyContext, we are creating our own context at the root level. So if root level context is cancelled, the children level contexts will also get cancel.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()
	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}
