package models

type Relationship struct {
	Id            int64 `json:"id"`
	FirstEmailId  int64 `json:"firstEmailId"`
	SecondEmailId int64 `json:"secondEmailId"`
	Status        int64 `json:"status"`
}
