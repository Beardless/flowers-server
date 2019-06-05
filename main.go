package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type envVariables struct {
	host   string
	port   int
	user   string
	dbname string
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/flower", getLastWatering).Methods("GET")
	r.HandleFunc("/flower", createNewWatering).Methods("POST")
	return r
}

func loadAndReturnEnvs() envVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	parsedPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Couldnt parse port number")
	}

	envs := envVariables{
		host:   os.Getenv("HOST"),
		port:   parsedPort,
		user:   os.Getenv("USER"),
		dbname: os.Getenv("DBNAME"),
	}

	return envs
}

func main() {
	envs := loadAndReturnEnvs()
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		envs.host, envs.port, envs.user, envs.dbname)

	db, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("PORT:", os.Getenv("PORT"))

	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 31, "jon@calhouaasdn.io", "Jonathansss", "Calhounsss").Scan(&id)
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
