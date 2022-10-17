package assignment

import (
	"sql_filler/subjects"

	"github.com/samonzeweb/godb"
)

func (json *Json) AddToDB(db *godb.DB, replace ...string) error {
	_, err := subjects.AddResourceDB(db, json, replace...)
	return err
}
