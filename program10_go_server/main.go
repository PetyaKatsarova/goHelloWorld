package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(res, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "\nName = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)
}

func helloHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)


	/*
	You start the server in a new goroutine so that the main goroutine can continue to listen for operating system signals.
	This non-blocking call is essential for the graceful shutdown mechanism.
	*/
	srv := &http.Server{Addr: ":4242"}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("Starting server at port 4242\n")
	// You set up a channel to listen for interruption signals (like Ctrl+C from the keyboard). The signal.Notify function configures the program to send signal notifications to the stop channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// block until a singal is received
	<-stop
	// create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// doesnt block if no connection, but will otherwise wait unit the ....
	srv.Shutdown(ctx)

	log.Println("Shutting down gracefully, press Ctrl+C again to force")


	// if err := http.ListenAndServe(":4242", nil); err != nil {
	// 	log.Fatal(err)
	// }
}