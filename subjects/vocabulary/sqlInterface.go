package vocabulary

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

func SelctVocabulary(db *godb.DB, id int) (vocabulary *Json, err error) {
	vocabulary = &Json{}
	err = common.GetSubjectDB(db, vocabulary)
	if err != nil {
		return
	}

	err = vocabulary.getAuxularyData(db)

	return
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
	_, err := db.DeleteFrom(common.ReadingTable).WhereQ(
		godb.Q(common.ReadingsIdRow+" = ?", json.ID),
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

func (json *Json) getAuxularyData(db *godb.DB) error {
	readings := make([]common.Readings, 0)
	err := db.Select(&readings).Where(common.SubjectIdRow+" = ?", json.ID).
		OrderBy("is_primary DESC").Do()
	if err != nil {
		return err
	}
	json.Data.Readings = readings

	contextSentences := make([]ContextSentences, 0)
	err = db.Select(&contextSentences).Where(ContextSentencesID+" = ?", json.ID).Do()
	if err != nil {
		return err
	}
	json.Data.ContextSentences = contextSentences

	pronunciationAudios := make([]PronunciationAudios, 0)
	err = db.Select(&pronunciationAudios).Where(PronunciationAudiosId+" = ?", json.ID).Do()
	if err != nil {
		return err
	}
	json.Data.PronunciationAudios = pronunciationAudios

	partsOfSpeech := make([]string, 0)
	err = db.SelectFrom(PartsOfSpeechTable).Columns("part_of_speech").
		Where(PartsOfSpeechId+" = ?", json.ID).Do(&partsOfSpeech)
	if err != nil {
		return err
	}
	json.Data.PartsOfSpeech = partsOfSpeech

	return nil
}
