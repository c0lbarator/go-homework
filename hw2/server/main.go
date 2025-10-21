package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hw2/api"
)

func versionHandler(writer http.ResponseWriter, request *http.Request) {
	resp := api.VersionResponse{Version: "v1.0.0"}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resp)
}

func decodeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req api.DecodeRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "bad request", http.StatusBadRequest)
		return
	}
	data, err := base64.StdEncoding.DecodeString(req.InputString)
	if err != nil {
		http.Error(writer, "invalid base64", http.StatusBadRequest)
		return
	}
	resp := api.DecodeResponse{OutputString: string(data)}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resp)
}

func hardOpHandler(writer http.ResponseWriter, request *http.Request) {
	min := 10
	max := 20
	delay := time.Duration(rand.Intn(max-min+1)+min) * time.Second
	select {
	case <-time.After(delay):
	case <-request.Context().Done():
		http.Error(writer, "request canceled", http.StatusRequestTimeout)
		return
	}
	if rand.Intn(2) == 0 {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, "ok")
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, "internal error")
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/hard-op", hardOpHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting server...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()

	<-done
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %v", err)
	}
	log.Println("server stopped")
}
