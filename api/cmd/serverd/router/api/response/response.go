package response

import (
	"encoding/json"
	"net/http"
)

//// Response stores info to retrieve response json
type Response struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type ErrorResp struct {
	ErrorMessage string
	Status       int
}

type SuccessResp struct {
	Success bool
}

func ResponseJson(w http.ResponseWriter, r *http.Request, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status int
	switch werr := object.(type) {
	case *ErrorResp:
		respBytes, _ = json.Marshal(werr)
		status = werr.Status
	default:
		status = http.StatusOK
	}

	w.WriteHeader(status)
	w.Write(respBytes)
}
