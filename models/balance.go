package models

import "time"

type Balance struct {
	Group     string             `bson:"group" json:"group"`
	Balances  map[string]float64 `bson:"balances" json:"balances"`
	TimeStamp time.Time          `bson:"updated_at" json:"updated_at"`
}
