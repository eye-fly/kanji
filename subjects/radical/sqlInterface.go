package radical

import (
	"sql_filler/subjects"

	"github.com/samonzeweb/godb"
)

func (json *Json) AddToDB(db *godb.DB, replace ...string) error {
	status, err := subjects.AddToDB(db, json, replace...)
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
	}
	return nil
}

func (json *Json) addAuxularyData(db *godb.DB) error {
	for _, img := range json.Data.CharacterImages {
		img.RadicalId = json.ID
		err := db.Insert(&img).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

func (json *Json) deleteAuxularyData(db *godb.DB) error {
	_, err := db.DeleteFrom(characterImagesTable).WhereQ(
		godb.Q(characterImagesIdRow+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	return nil
}
