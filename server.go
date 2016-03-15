package main
import (
	"github.com/gorilla/mux"
	"summary-api/summary"
	"log"
	"net/http"
	"encoding/json"
	)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/summary", Summary).Methods("POST")

	log.Print(http.ListenAndServe(":8080", router))
}

func Summary(w http.ResponseWriter, r *http.Request) {
	a := &summary.Article{}
	err := json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		panic(err)
	}

	s, e := json.Marshal(a.Summary())
	if e != nil {
		panic(e)
	}
	        w.Header().Set("Content-Type", "application/json")

	w.Write(s)
}
