package application

import (
	"context"
	"fmt"
	"net/http"

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
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping().Err()
	if err != nil {
		return fmt.Errorf("failed to start the redis : %w", err)
	}
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
		return server.Shutdown(ctx)
	}

	return nil

}
