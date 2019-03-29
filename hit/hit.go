package hit

import (
	"net/url"
	"strconv"
)

//Hit represents an object store hit in memory
type Hit struct {
	higherHit           int
	mostFrequentRequest string
	requests            map[string]int
}

//New initialize a Hit object
func New() *Hit {
	h := &Hit{}
	h.requests = make(map[string]int)
	return h
}

//Add is used to add a request in Hit
func (h *Hit) Add(int1, int2, limit int, str1, str2 string) {
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
}

//GetMostFrequentRequest return the most frequebt request
func (h *Hit) GetMostFrequentRequest() string {
	return h.mostFrequentRequest
}

func makeKey(int1, int2, limit int, str1, str2 string) string {
	q := url.Values{}
	q.Add("int1", strconv.Itoa(int1))
	q.Add("int2", strconv.Itoa(int2))
	q.Add("limit", strconv.Itoa(limit))
	q.Add("str1", str1)
	q.Add("str2", str2)
	return q.Encode()
}
