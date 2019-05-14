// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"encoding/json"
	"fmt"
	"io/ioutil"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/flower", getLastWatering).Methods("GET")
	r.HandleFunc("/flower", createNewWatering).Methods("POST")
	return r
}

func main() {
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

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
