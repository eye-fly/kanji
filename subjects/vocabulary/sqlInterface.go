package vocabulary

import (
	"database/sql"
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

func SelectVocabulary(db *godb.DB, id int) (vocabulary *Json, err error) {
	vocabulary = &Json{ID: id}
	err = common.GetSubjectDB(db, vocabulary)
	if err != nil {
		return
	}

	err = vocabulary.getAuxularyData(db)

	return
}

func FindVocabulary(db *godb.DB, characters string) (vocabulary *Json, err error) {
	vocabulary = &Json{}

	err = db.Select(vocabulary).Where(vacabularyCharactersRow+" = ?", characters).Do()
	if err != nil {
		return
	}
	err = common.GetSubjectDB(db, vocabulary)
	if err != nil {
		return
	}

	err = vocabulary.getAuxularyData(db)

	return
}

func FindVocabularyByReading(db *godb.DB, characters string) (vocs []*Json, err error) {
	readings := make([]common.Readings, 0)
	err = db.Select(&readings).Where(common.ReadingReadingRow+" = ?", characters).Do()
	if err != nil {
		return nil, err
	}

	s := map[int]bool{}
	for _, reading := range readings {
		s[reading.SubjectId] = true
	}
	vocs = make([]*Json, 0)
	for id, _ := range s {
		vc, err := SelectVocabulary(db, id)
		if err != sql.ErrNoRows {
			if err != nil {
				return nil, err
			}
			vocs = append(vocs, vc)
		}
	}
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
	err := db.Select(&readings).Where(common.ReadingsIdRow+" = ?", json.ID).
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

	type partOfSpeech struct {
		PartOfSpeach string `db:"part_of_speech"`
	}
	partsOfSpeech := make([]partOfSpeech, 0)
	err = db.SelectFrom(PartsOfSpeechTable).Columns("part_of_speech").
		Where(PartsOfSpeechId+" = ?", json.ID).Do(&partsOfSpeech)
	if err != nil {
		return err
	}
	json.Data.PartsOfSpeech = make([]string, len(partsOfSpeech))
	for i, v := range partsOfSpeech {
		json.Data.PartsOfSpeech[i] = v.PartOfSpeach
	}

	return nil
}
