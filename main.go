package main

import (
	"fmt"
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

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MyMiddleware")
		w.Write([]byte("mymiddle"))
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(myMiddleware)

	r.Get("/", http.NotFoundHandler().ServeHTTP)
	r.Get("/hello", helloHandler)
	r.Get(`/users/{id}`, userHandler)

	http.ListenAndServe(":8080", r)
	fmt.Println("Start Server")

}
