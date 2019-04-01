package hit

//Hit represents an object store hit in memory
type Hit struct {
	higherHit           float64
	mostFrequentRequest string
	requests            map[string]float64
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

	if _, exists := h.requests[key]; exists {
		h.requests[key]++
	} else {
		h.requests[key] = 1
	}

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
