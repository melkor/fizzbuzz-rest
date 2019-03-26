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
	paramsIntVal := make(map[string]int)

	for _, name := range []string{"int1", "int2", "limit"} {
		val, err := strconv.Atoi(params[name])
		if err != nil {
			w.WriteHeader(400)
			log.Println("query error : ", err)
			json.NewEncoder(w).Encode(map[string]string{"error": name + " must be an integer"})
			return
		}
		paramsIntVal[name] = val
	}

	result, err := fizzbuzz.Fizzbuzz(paramsIntVal["int1"], paramsIntVal["int2"], paramsIntVal["limit"], params["str1"], params["str2"])

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
