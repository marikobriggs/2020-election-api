package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// inspo https://www.soberkoder.com/go-rest-api-mysql-gorm/

type Result struct {
	ID           int    `json:"id"`
	State        string `json:"state"`
	EC           int    `json:"ec"`
	TrumpPercent int    `json:"trumpPercent"`
	BidenPercent int    `json:"bidenPercent"`
}

type Results []Result

func getResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputResultID := params["id"]

	var result Result
	db.Preload("Results").First(&result, inputResultID)
	json.NewEncoder(w).Encode(result)
}

func getResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []Result

	// Find() fetches all results
	db.Preload("Result").Find(&results)

	fmt.Println("allResults endpoint hit")
	json.NewEncoder(w).Encode(results)

}

func postResults(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// not a great use case for crud :)
	myRouter.HandleFunc("/", homePage)
	// create
	myRouter.HandleFunc("/results", postResults).Methods("POST")
	// read
	myRouter.HandleFunc("/results", getResult).Methods("GET")
	// read all
	myRouter.HandleFunc("/results", getResults).Methods("GET")

	initDB()

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

var db *gorm.DB

func initDB() {
	var err error
	// dataSourceName := "your_username:your_password@tcp(localhost:3306)/your_database_name?parseTime=True"
	dataSource := "root:@tcp(localhost:8080)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSource)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//create db - comment out if db is already created
	db.Exec("CREATE DATABASE resultsdb")
	db.Exec("USE resultsdb")
}

func main() {
	handleRequests()
}
