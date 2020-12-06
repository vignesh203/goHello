package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goHello/handler"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	// r := mux.NewRouter()
	// //r.Host("www.localhost").Path("/").HandlerFunc(Handler)
	// r.HandleFunc("/", Handler)
	// err := http.ListenAndServe(":9000", r)
	// if err != nil {
	// 	log.Fatal("ListenAndServe error: ", err)
	// }
	// Parsing wait for graceful shutdown setting
	wait, err := time.ParseDuration("3s")

	if err != nil {
		log.Printf("Failed to parse wait duration of '%s' shutdown without waiting for connections to close.", "3s")
		wait = 0
	}

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Session-Key", "Device-ID"},
		Debug:          true,
	})

	//itemAPI := handler.NewItemAPI()

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", handler.OnPing).Methods("GET")

	handler := cor.Handler(router)
	address := fmt.Sprintf(":%s", "80")

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Started server listening at %s", address)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	log.Printf("Waiting %s for connections to close.", wait.String())

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Server gracefully shutdown.")
	os.Exit(0)
}
