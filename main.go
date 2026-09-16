package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // string
	sms := "user " + id
	w.Write([]byte(sms))

}

func main() {
	r := chi.NewRouter()

	r.Get("/", http.NotFoundHandler().ServeHTTP)
	r.Get("/hello", helloHandler)
	r.Get(`/users/{id}`, userHandler)

	http.ListenAndServe(":8080", r)

}
