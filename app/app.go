package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/melkor/fizzbuzz-rest/fizzbuzz"
	log "github.com/sirupsen/logrus"
)

//App struct used to handle application
type App struct {
	address string
	router  *mux.Router
}

//Init initialize application
func New(address string) *App {

	// assume http.ListenAndServe check if address is valide
	if address == "" {
		address = ":8000"
	}

	a := &App{
		address: address,
		router:  mux.NewRouter(),
	}

	return a
}

//GetFizzBuzz is the endpoint behind the route "/fizzbuzz"
// which will call fizzbuzz func and return result as a json message
func (a *App) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	log.Debugln("received a fizzbuzz request")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	paramsIntVal := make(map[string]int)

	//convert into integer the parameters from querystring
	// which must be handle as integer
	for _, name := range []string{"int1", "int2", "limit"} {
		val, err := strconv.Atoi(params[name])
		log.Debugln("	param: ", name, ", value: ", val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorln("query error : ", err)
			json.NewEncoder(w).Encode(map[string]string{"error": name + " must be an integer"})
			return
		}
		paramsIntVal[name] = val
	}

	log.Debugln("	param: str1, value: ", params["str1"])
	log.Debugln("	param: str1, value: ", params["str2"])

	// call fizz-buzz function
	result, err := fizzbuzz.Fizzbuzz(
		paramsIntVal["int1"],
		paramsIntVal["int2"],
		paramsIntVal["limit"],
		params["str1"],
		params["str2"],
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorln("fizzbuzz error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	log.Debugln("	result: ", result)

	json.NewEncoder(w).Encode(result)
	log.Debugln("Done...")
}

//add http route into this func
func (a *App) initializeRoutes() {
	log.Infoln("Initialize routes")

	//route "GET" to fizzbuzzaddr
	log.Debugln("Add route to /fizzbuzz endpoint")
	a.router.HandleFunc("/fizzbuzz", a.GetFizzBuzz).
		Methods("GET").
		Queries("int1", "{int1}").
		Queries("int2", "{int2}").
		Queries("limit", "{limit}").
		Queries("str1", "{str1}").
		Queries("str2", "{str2}")

	//TODO add endpoint to stats
}

//Run function launch application
func (a *App) Run() {
	log.Infoln("Launch application")
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(a.address, a.router))
}
