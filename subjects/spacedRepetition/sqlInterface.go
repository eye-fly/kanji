package spaced_repetition

import (
	"sql_filler/subjects/common"

	"github.com/samonzeweb/godb"
)

func (json *Json) AddToDB(db *godb.DB, replace ...string) error {
	status, err := common.AddResourceDB(db, json, replace...)
	if status == common.AddStatus {
		err = json.addAux(db)
		if err != nil {
			return err
		}
	} else if status == common.ReplaceStatus {
		json.replaceAux(db)
		if err != nil {
			return err
		}
	}
	return err
}

func (json *Json) addAux(db *godb.DB) error {

	for _, stage := range json.Data.Stages {
		stage.SrsId = json.ID
		err := db.Insert(&stage).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

func (json *Json) replaceAux(db *godb.DB) error {
	_, err := db.DeleteFrom(json.Data.Stages[0].TableName()).WhereQ(
		godb.Q(SrsSrageIdRow+" = ?", json.GetId()),
	).Do()
	if err != nil {
		return err
	}
	return json.addAux(db)
}
