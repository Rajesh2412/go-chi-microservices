package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(callHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to run the server", err)
	}

}

func callHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
