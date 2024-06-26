package kanji

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

func SelctKanji(db *godb.DB, id int) (kanji *Json, err error) {
	kanji = &Json{
		ID: id,
	}
	err = common.GetSubjectDB(db, kanji)
	if err != nil {
		return
	}

	err = kanji.getAuxularyData(db)

	return
}

func FindKanji(db *godb.DB, char string) (kanji *Json, err error) {
	kanji = &Json{}
	err = db.Select(kanji).Where(kanjiCharactersRow+" = ?", char).Do()
	if err != nil {
		return
	}

	return SelctKanji(db, kanji.ID)
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
	_, err := db.DeleteFrom(common.ReadingTable).WhereQ(
		godb.Q(common.ReadingsIdRow+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	return nil
}

func (json *Json) getAuxularyData(db *godb.DB) error {
	readings := make([]common.Readings, 0)
	err := db.Select(&readings).Where(common.ReadingsIdRow+" = ?", json.ID).
		OrderBy("is_primary DESC").Do()
	if err != nil {
		return err
	}

	json.Data.Readings = readings

	return nil
}
