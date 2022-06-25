package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func dflt(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func nameParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func data(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func headers(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		panic(err)
	}

	w.Header().Add("a+b", strconv.Itoa(a+b))
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", nameParam).Methods("GET")
	router.HandleFunc("/bad", bad).Methods("GET")
	router.HandleFunc("/data", data).Methods("POST")
	router.HandleFunc("/headers", headers).Methods("POST")
	router.PathPrefix("/").HandlerFunc(dflt)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
