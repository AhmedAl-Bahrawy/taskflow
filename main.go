package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	_ = pgx.ErrNoRows

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "TaskFlow API is running")
	})

	fmt.Println("TaskFlow API starting on :8080")
	http.ListenAndServe(":8080", r)
}