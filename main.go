package main

import (
	"encoding/json"
	"flowers-server/config"
	"flowers-server/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/flower", getWaterings).Methods("GET")
	r.HandleFunc("/flower", createNewWatering).Methods("POST")
	return r
}

func main() {
	envs := config.ReturnEnvs()
	fmt.Println(envs)
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		envs.Host, envs.Port, envs.User, envs.Dbname)
	config.InitDB(psqlConnString)

	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func getWaterings(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	waterings, error := models.GetAllWaterings()
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, err := json.Marshal(waterings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}

func createNewWatering(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	watering, err := models.CreateNewWatering(body)
	js, err := json.Marshal(watering)

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
