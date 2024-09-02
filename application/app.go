package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {

	app := &App{
		router: loadRoutes(),
	}
	fmt.Printf("this is type of app %T", app)
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	fmt.Printf("this is context pass value %s", ctx.Value("mykey"))

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("Failed to start the server %w", err)
	}

	return nil

}
