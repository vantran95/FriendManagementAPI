package models

// Relationship stores info to retrieve relationship struct
type Relationship struct {
	ID        int64  `json:"id"`
	RequestID int64  `json:"request_id"`
	TargetID  int64  `json:"target_id"`
	Status    string `json:"status"`
}
