package hit

//This package is a mock of Cache (see redis.go) used in units test of the
//application (app_test.go).

//Hit represents an object store hit in memory
type Hit struct {
	higherHit           float64            //the score of the mostFrequentRequest
	mostFrequentRequest string             //the most frequent request
	requests            map[string]float64 //historic of all requests
}

//NewMock initialize a Hit mocker object
func NewMock() *Hit {
	h := &Hit{}
	h.requests = make(map[string]float64)
	return h
}

//Add is used to add a request in Hit
func (h *Hit) Add(int1, int2, limit int, str1, str2 string) (float64, error) {
	key := makeKey(int1, int2, limit, str1, str2)

	// if request is not present in historic, we initialize the score of if to 0
	// else we add 1 to that score
	if _, exists := h.requests[key]; exists {
		h.requests[key]++
	} else {
		h.requests[key] = 1
	}

	// if score it higer thant higherHit, this request become the most
	// frequent request
	if h.requests[key] > h.higherHit {
		h.higherHit = h.requests[key]
		h.mostFrequentRequest = key
	}

	return h.requests[key], nil
}

//GetMostFrequentRequest return the most frequebt request
func (h *Hit) GetMostFrequentRequest() (string, error) {
	return h.mostFrequentRequest, nil
}
