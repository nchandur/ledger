package models

import "time"

type Group struct {
	TimeStamp time.Time `bson:"created_at" json:"created_at"`
	GroupName string    `bson:"group_name" json:"group_name"`
	People    []string  `bson:"people" json:"people"`
	Currency  string    `bson:"currency" json:"currency"`
}
