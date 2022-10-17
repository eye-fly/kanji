package vocabulary

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

	for _, sentence := range json.Data.ContextSentences {
		sentence.VocabularyId = json.ID
		err := db.Insert(&sentence).Do()
		if err != nil {
			return err
		}
	}
	for _, audio := range json.Data.PronunciationAudios {
		audio.VocabularyId = json.ID
		err := db.Insert(&audio).Do()
		if err != nil {
			return err
		}
	}
	for _, partOfSpeach := range json.Data.PartsOfSpeech {
		part := PartsOfSpeach{
			VocabularyId: json.ID,
			PartOfSpeach: partOfSpeach,
		}
		err := db.Insert(&part).Do()
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

	_, err = db.DeleteFrom(ContextSentencesTable).WhereQ(
		godb.Q(ContextSentencesID+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	_, err = db.DeleteFrom(PronunciationAudiosTable).WhereQ(
		godb.Q(PronunciationAudiosId+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}
	_, err = db.DeleteFrom(PartsOfSpeechTable).WhereQ(
		godb.Q(PartsOfSpeechId+" = ?", json.ID),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
