package models

import "time"

type People struct {
	Id          uint      `json:"id"`
	Createdat   time.Time `json:"created_at"`
	Nome        string    `json:"nome"`
	Sexualidade string    `json:"sexualidade"`
}
