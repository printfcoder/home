package domain

import "time"

type Tag struct {
	Name string `json:"name"`
}

type Expense struct {
	Tags   []Tag     `json:"tags"`
	Time   time.Time `json:"time"`
	Remark string    `json:"remark"`
}
