package models

type Relationship struct {
	Id            int64 `json:"id"`
	FirstEmailId  int64 `json:"firstEmailId"`
	SecondEmailID int64 `json:"secondEmailID"`
	Status        int64 `json:"status"`
}
