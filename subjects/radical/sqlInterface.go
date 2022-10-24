package radical

import (
	"sql_filler/subjects/common"

	"github.com/samonzeweb/godb"
)

func (json *Json) AddToDB(db *godb.DB, replace ...string) error {
	_, err := common.AddSubjectDB(db, json, replace...)
	if err != nil {
		return err
	}

	err = json.deleteAuxularyData(db)
	if err != nil {
		return err
	}
	err = json.addAuxularyData(db)
	if err != nil {
		return err
	}
	return nil
}

func SelctRadical(db *godb.DB, id int) (radical *Json, err error) {
	radical = &Json{
		ID: id,
	}
	err = common.GetSubjectDB(db, radical)

	return
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
		godb.Q(CharacterImagesIdRow+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	return nil
}
