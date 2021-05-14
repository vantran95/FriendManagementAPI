package v1

// Response stores info to retrieve response json
type Response struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
