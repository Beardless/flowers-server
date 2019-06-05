package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "bartlomiejmikolajczuk"
	dbname = "bartlomiejmikolajczuk"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/flower", getLastWatering).Methods("GET")
	r.HandleFunc("/flower", createNewWatering).Methods("POST")
	return r
}

func main() {
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	r := newRouter()
	http.ListenAndServe(":8080", r)
}

// Watering type
type Watering struct {
	ID          string `json:"id"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
}

var waterings = []Watering{
	Watering{
		ID:          "my test first watering",
		Timestamp:   "My first test timestamp",
		Description: "Description of watering",
	},
}

func getLastWatering(w http.ResponseWriter, r *http.Request) {
	lastWatering := waterings[len(waterings)-1]

	wateringList, err := json.Marshal(lastWatering)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(wateringList)
}

func createNewWatering(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var watering Watering
	err = json.Unmarshal(b, &watering)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	waterings = append(waterings, watering)

	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
