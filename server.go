package main

import (
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "I'm healthy", http.StatusOK)
}

func init() {

	http.Handle("/_ah/health", http.HandlerFunc(healthCheckHandler))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Hi.", http.StatusOK)
	}))

}
