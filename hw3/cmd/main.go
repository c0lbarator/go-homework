package main

import (
	"context"
	"hw3/internal/application/usecase"
	"hw3/internal/infrastructure/api/balance"
	"hw3/internal/infrastructure/api/register"
	"hw3/internal/infrastructure/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	storage := &db.InmemoryStorage{}
	storage.CreateRepository()
	balanceUsecase := usecase.BalanceUsecase{
		AccountRepo: storage,
	}
	registerUsecase := usecase.RegisterUsecase{
		AccountRepo: storage,
	}
	balanceHandler := balance.NewHandler(balanceUsecase)
	registerHandler := register.NewHandler(registerUsecase)
	mux := http.NewServeMux()
	mux.HandleFunc("/balance", balanceHandler.BalanceHandler)
	mux.HandleFunc("/deposit", balanceHandler.DepositHandler)
	mux.HandleFunc("/transfer", balanceHandler.TransferHandler)
	mux.HandleFunc("/register", registerHandler.RegisterHandler)
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
