package domain

import "time"

type Orders struct {
	Id        int       `json:"id" gorm:"primary_key, auto_increment"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
