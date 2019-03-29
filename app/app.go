package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/melkor/fizzbuzz-rest/fizzbuzz"
	"github.com/melkor/fizzbuzz-rest/hit"
	log "github.com/sirupsen/logrus"
)

//App struct used to handle application
type App struct {
	address string
	hit     *hit.Cache
	router  *mux.Router
}

//New initialize application
func New(listenAddress, cacheAddress, cachePassword string) *App {

	//only check if an address was set, the validity check
	// will be done by http.ListenAndServ
	if listenAddress == "" {
		listenAddress = ":8000"
	}

	if cacheAddress == "" {
		cacheAddress = "localhost:6379"
	}

	a := &App{
		address: listenAddress,
		hit:     hit.NewCache(cacheAddress, cachePassword, 0),
		router:  mux.NewRouter(),
	}

	return a
}

//getFizzBuzz is the endpoint behind the route "/fizzbuzz"
// which will call fizzbuzz func and return result as a json message
func (a *App) getFizzBuzz(w http.ResponseWriter, r *http.Request) {
	log.Debugln("received a fizzbuzz request")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	paramsIntVal := make(map[string]int)

	// this function is called only if all parameters are set in
	// the querystring, we don't need to check if keys (int1 ,int2,
	// limit, str1 and str2) are present in params map

	//convert into integer the parameters from querystring
	// which must be handled as integer
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
	log.Debugln("	param: str2, value: ", params["str2"])

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

	log.Debugln("Add request into historic")
	score, err := a.hit.Add(
		paramsIntVal["int1"],
		paramsIntVal["int2"],
		paramsIntVal["limit"],
		params["str1"],
		params["str2"],
	)

	if err != nil {
		log.Errorln("add hit error : ", err)
	}

	log.Debugln("hit score: ", score)
}

func (a *App) getMostFrequentRequest(w http.ResponseWriter, r *http.Request) {
	log.Debugln("received an audience request")

	w.Header().Set("Content-Type", "application/json")

	result, err := a.hit.GetMostFrequentRequest()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln("most frequent request error : ", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(result)

	log.Debugln("	result: ", result)
}

//add http route into this func
func (a *App) initializeRoutes() {
	log.Infoln("Initialize routes")

	//route "GET" to fizzbuzzaddr
	//Note: all parameters must be set to use this route
	log.Debugln("Add route to /fizzbuzz endpoint")
	a.router.HandleFunc("/fizzbuzz", a.getFizzBuzz).
		Methods("GET").
		Queries("int1", "{int1}").
		Queries("int2", "{int2}").
		Queries("limit", "{limit}").
		Queries("str1", "{str1}").
		Queries("str2", "{str2}")

	//endpoint to most frequent request
	log.Debugln("Add route to /mostFrequentRequest endpoint")
	a.router.HandleFunc("/mostFrequentRequest", a.getMostFrequentRequest)

}

//Run function launch application
func (a *App) Run() {
	log.Infoln("Launch application")
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(a.address, a.router))
}
