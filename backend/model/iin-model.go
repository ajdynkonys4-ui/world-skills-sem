package model

import "time"

type IINModel struct {
	ID         int       `json:"id"`
	Iin        string    `json:"iin"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"time"`
}
