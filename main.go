package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte("hello"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // string
	sms := "user " + id
	w.Write([]byte(sms))

}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metod := r.Method
		path := r.URL.Path
		fmt.Println("начало отчета")
		t := time.Now()

		next.ServeHTTP(w, r)

		endtime := time.Since(t)

		fmt.Println(metod + " " + path + " " + endtime.String())
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("second До")
		next.ServeHTTP(w, r)
		fmt.Println("second после")
	})
}

func main() {
	r := chi.NewRouter()
	//r.Use(secondMiddleware)
	r.Use(myMiddleware)

	r.Get("/", http.NotFoundHandler().ServeHTTP)
	r.Get("/hello", helloHandler)
	r.Get(`/users/{id}`, userHandler)

	http.ListenAndServe(":8080", r)

}
