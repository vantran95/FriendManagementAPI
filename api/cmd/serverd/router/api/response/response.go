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

// Error stores info to retrieve error response json
type Error struct {
	Status      int
	Code        string
	Description string
}

type Result struct {
	Success bool `json:"success"`
}

func ResponseJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status int
	switch werr := object.(type) {
	case Error:
		respBytes, _ = json.Marshal(werr)
		status = werr.Status
	default:
		status = http.StatusOK
	}

	w.WriteHeader(status)
	w.Write(respBytes)
}
