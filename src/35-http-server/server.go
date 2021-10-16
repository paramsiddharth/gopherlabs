package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}
		fmt.Fprintf(rw, "Hello, %s!\n", strings.Title(strings.ToLower(name)))
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		switch <-c {
		case os.Interrupt:
			fmt.Println("\rExitting...")
			os.Exit(0)
		}
	}()

	fmt.Println("Listening on http://localhost:8080.")
	http.ListenAndServe(":8080", server)
}
