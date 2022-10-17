package kanji

import (
	"sql_filler/subjects"

	"github.com/samonzeweb/godb"
)

func (json *Json) AddToDB(db *godb.DB, replace ...string) error {
	status, err := subjects.AddSubjectDB(db, json, replace...)
	if err != nil {
		return err
	}

	if status == subjects.AddStatus {
		err = json.addAuxularyData(db)
		if err != nil {
			return err
		}
	} else if status == subjects.ReplaceStatus {
		err = json.deleteAuxularyData(db)
		if err != nil {
			return err
		}
		err = json.addAuxularyData(db)
		if err != nil {
			return err
		}
	}
	return nil
}

func (json *Json) addAuxularyData(db *godb.DB) error {
	for _, reading := range json.Data.Readings {
		reading.SubjectId = json.ID
		err := db.Insert(&reading).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

func (json *Json) deleteAuxularyData(db *godb.DB) error {
	_, err := db.DeleteFrom(subjects.ReadingTable).WhereQ(
		godb.Q(subjects.ReadingsIdRow+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	return nil
}
