package main

import (
	"encoding/json"
	"fizzbuzz-rest/fizzbuzz"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type parameters struct {
	int1  int
	int2  int
	limit int
	str1  string
	str2  string
}

func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//TODO handle error
	int1, err := strconv.Atoi(params["int1"])
	if err != nil {
		w.WriteHeader(400)
		log.Println("query error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "int1 must be an integer"})
		return
	}

	int2, err := strconv.Atoi(params["int2"])
	if err != nil {
		w.WriteHeader(400)
		log.Println("query error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "int2 must be an integer"})
		return
	}

	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		w.WriteHeader(400)
		log.Println("query error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "limit must be an integer"})
		return
	}

	str1 := params["str1"]
	str2 := params["str2"]

	result, err := fizzbuzz.Fizzbuzz(int1, int2, limit, str1, str2)

	if err != nil {
		w.WriteHeader(400)
		log.Println("fizzbuzz error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", GetFizzBuzz).
		Methods("GET").
		Queries("int1", "{int1}").
		Queries("int2", "{int2}").
		Queries("limit", "{limit}").
		Queries("str1", "{str1}").
		Queries("str2", "{str2}")

	log.Fatal(http.ListenAndServe(":8000", router))

}
