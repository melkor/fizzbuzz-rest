package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFizzBuzz(t *testing.T) {
	testCase := []struct {
		label          string
		int1           string
		int2           string
		limit          string
		str1           string
		str2           string
		expectedStatus int
		expectedBody   string
	}{
		{
			label:          "wrong parameter: int1 is empty",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"error\":\"int1 must be an integer\"}\n",
		},
		{
			label:          "wrong parameter: int1 not a valid integer",
			int1:           "fizzbuzz",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"error\":\"int1 must be an integer\"}\n",
		},
		{
			label:          "wrong parameter: int1 == 0",
			int1:           "0",
			int2:           "2",
			limit:          "3",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"error\":\"int1 can't be lower than 1\"}\n",
		},
		{
			label:          "wrong parameter: int1 == 0",
			int1:           "2",
			int2:           "3",
			limit:          "10",
			str1:           "fizz",
			str2:           "buzz",
			expectedStatus: http.StatusOK,
			expectedBody:   "[\"1\",\"fizz\",\"buzz\",\"fizz\",\"5\",\"fizzbuzz\",\"7\",\"fizz\",\"buzz\",\"fizz\"]\n",
		},
	}
	for _, tc := range testCase {
		req, err := http.NewRequest("GET", "/fizzbuzz", nil)
		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()

		q.Add("int1", tc.int1)
		q.Add("int2", tc.int2)
		q.Add("limit", tc.limit)
		q.Add("str1", tc.str1)
		q.Add("str2", tc.str2)

		req.URL.RawQuery = q.Encode()

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		app := New("", "", "")
		app.initializeRoutes()

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		app.router.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tc.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expectedStatus)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != tc.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), tc.expectedBody)
		}

	}
}
