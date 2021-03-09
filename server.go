package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Result struct {
	ID           int    `json:"id"`
	State        string `json:"state"`
	EC           int    `json:"ec"`
	TrumpPercent int    `json:"trumpPercent"`
	BidenPercent int    `json:"bidenPercent"`
}

type Results []Result

func allResults(w http.ResponseWriter, r *http.Request) {
	results := Results{
		Result{ID: 99, State: "test state", EC: 10, TrumpPercent: 48, BidenPercent: 50},
	}

	fmt.Println("allResults endpoint hit")
	json.NewEncoder(w).Encode(results)

}

func testPostResults(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/results", allResults).Methods("GET")
	myRouter.HandleFunc("/results", testPostResults).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
	handleRequests()
}
