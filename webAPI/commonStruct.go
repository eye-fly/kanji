package webapi

import "github.com/samonzeweb/godb"

type BackEnd struct {
	db *godb.DB
}

func NewBackEnd(db *godb.DB) *BackEnd {
	return &BackEnd{
		db: db,
	}
}

type AuxiliaryReadings struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Reading string `json:"reading"`
}
