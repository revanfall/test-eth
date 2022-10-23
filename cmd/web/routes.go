package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"test-eth/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Transactions)
	mux.Get("/p={pnum}&lim={lnum}", handlers.Repo.TransactionsPaginationWithLimit)
	mux.Get("/p={pnum}", handlers.Repo.TransactionsPagination)
	mux.Get("/hash={hash}", handlers.Repo.TransactionByHash)
	mux.Get("/s={s}", handlers.Repo.TransactionBySender)
	mux.Get("/r={r}", handlers.Repo.TransactionByReceiver)
	mux.Get("/timestamp={ts}", handlers.Repo.TransactionsByTimeStamp)

	return mux
}
