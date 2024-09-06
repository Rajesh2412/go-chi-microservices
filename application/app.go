package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

// this struct is to include required dependents for our application.
type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {

	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := a.rdb.Ping().Err()
	if err != nil {
		return fmt.Errorf("failed to start the redis : %w", err)
	}
	//this is to close the redis connection
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis connection %w", err)
		}
	}()
	fmt.Println("Starting Server")

	ch := make(chan error, 1) // this channel is provide communication between goroutines below function
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start the server: %w", err)
		}
		close(ch) // closing the channel if the server is shitdown due to client offline or someother reason
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)

		defer cancel()

		return server.Shutdown(timeout)
	}

}
