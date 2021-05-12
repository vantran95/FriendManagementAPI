package models

// Relationship stores info to retrieve relationship struct
type Relationship struct {
	ID            int64 `json:"id"`
	FirstEmailID  int64 `json:"firstEmailID"`
	SecondEmailID int64 `json:"secondEmailID"`
	Status        int64 `json:"status"`
}
