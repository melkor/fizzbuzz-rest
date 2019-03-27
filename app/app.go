package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/melkor/fizzbuzz-rest/fizzbuzz"
)

//App struct used to handle application
type App struct {
	router *mux.Router
}

//Init initialize application
func (a *App) Init() {
	a.router = mux.NewRouter()
}

func (a *App) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
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

func (a *App) initializeRoutes() {
	//route "GET" to fizzbuzz
	a.router.HandleFunc("/fizzbuzz", a.GetFizzBuzz).
		Methods("GET").
		Queries("int1", "{int1}").
		Queries("int2", "{int2}").
		Queries("limit", "{limit}").
		Queries("str1", "{str1}").
		Queries("str2", "{str2}")
}

//Run function launch application
func (a *App) Run(addr string) {
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(addr, a.router))
}
