package main

import (
	"os"
	"net/http"
	"encoding/json"
	"summary-api/summary"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/summary", Summary).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8080", handlers.RecoveryHandler()(loggedRouter))
}

func Summary(w http.ResponseWriter, r *http.Request) {
	var (
		a = &summary.Article{}
		err error = nil
	)
	err = json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		panic(err)
	}

	s, err := json.Marshal(summary.Generate(a))
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(s)
}
