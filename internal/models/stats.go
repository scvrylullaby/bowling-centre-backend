package models

type Stats struct {
	Waiting  int `json:"waiting"`
	Playing  int `json:"playing"`
	Finished int `json:"finished"`
	Left     int `json:"left"`
}
