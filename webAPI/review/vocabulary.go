package review

import (
	"fmt"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/common"
	"sql_filler/subjects/vocabulary"
	webapi "sql_filler/webAPI"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type reviewVocabulary struct {
	En                []string                   `json:"en"`
	ID                int                        `json:"id"`
	Aud               []aud                      `json:"aud"`
	Voc               string                     `json:"voc"`
	Kana              []string                   `json:"kana"`
	Type              string                     `json:"type"`
	Kanji             []vacabularyKanji          `json:"kanji"`
	Characters        string                     `json:"characters"`
	AuxiliaryMeanings []common.AuxiliaryMeaning  `json:"auxiliary_meanings"`
	AuxiliaryReadings []webapi.AuxiliaryReadings `json:"auxiliary_readings"` //TODO
	Srs               int                        `json:"srs"`
	Syn               []interface{}              `json:"syn"`
}
type aud struct {
	URL           string `json:"url"`
	ContentType   string `json:"content_type"`
	Pronunciation string `json:"pronunciation"`
	VoiceActorID  int    `json:"voice_actor_id"`
}
type vacabularyKanji struct {
	En         string `json:"en"`
	ID         int    `json:"id"`
	Ja         string `json:"ja"`
	Kan        string `json:"kan"`
	Type       string `json:"type"`
	Characters string `json:"characters"`
}

func GetVocabulary(db *godb.DB, subjectID, userID int) (*reviewVocabulary, error) {
	var item vocabulary.Json
	err := db.Select(&item).Where("id = ?", subjectID).Do()
	if err != nil {
		return nil, err
	}
	reviewItem := startVoc()
	err = copier.Copy(&reviewItem, &item)
	if err != nil {
		return nil, fmt.Errorf("coping vocabulary: %w", err)
	}
	err = copier.Copy(&reviewItem, &item.Data)
	if err != nil {
		return nil, fmt.Errorf("coping vocabulary data: %w", err)
	}
	reviewItem.Voc = reviewItem.Characters
	reviewItem.Type = "Vocabulary"

	reviewItem.Srs, err = assignment.GetSrsStage(db, userID, subjectID)
	if err != nil {
		return nil, err
	}
	reviewItem.En, reviewItem.AuxiliaryMeanings, err = getAuxData(db, subjectID)
	if err != nil {
		return nil, fmt.Errorf("auxData vocabulary: %w", err)
	}

	err = reviewItem.getKana(db)
	if err != nil {
		return nil, err
	}
	err = reviewItem.getKanji(db, userID)
	if err != nil {
		return nil, err
	}
	err = reviewItem.getVoc(db)
	if err != nil {
		return nil, err
	}

	return &reviewItem, nil
}

func (voc *reviewVocabulary) getKana(db *godb.DB) error {
	readings := make([]common.Readings, 0)
	err := db.Select(&readings).Where(common.ReadingsIdRow+" = ?", voc.ID).
		Do()
	if err != nil {
		return err
	}
	for _, reading := range readings {
		if reading.AcceptedAnswer {
			voc.Kana = append(voc.Kana, reading.Reading)
		}
	}
	return nil
}

func (voc *reviewVocabulary) getKanji(db *godb.DB, userID int) error {
	relations := make([]common.AmalgamationSubjectIds, 0)
	err := db.Select(&relations).Where(common.AmalgamationSubjectSetRow+" = ?", voc.ID).
		Do()
	if err != nil {
		return err
	}
	for _, v := range relations {
		kanji, _ := GetKanji(db, v.ComponentId, userID)
		// if err != nil {
		// 	return err
		// } TODO: err if nece
		vocKanji, _ := toVocKanji(kanji)
		voc.Kanji = append(voc.Kanji, vocKanji)
	}
	return nil
}

func (voc *reviewVocabulary) getVoc(db *godb.DB) error {
	audios := make([]vocabulary.PronunciationAudios, 0)
	err := db.Select(&audios).Where(vocabulary.PronunciationAudiosId+" = ? AND content_type = \"audio/webm\"", voc.ID).
		Do()
	if err != nil {
		return err
	}

	var audio *aud
	for _, v := range audios {
		audio = &aud{}
		err = copier.Copy(&audio, &v)
		if err != nil {
			return fmt.Errorf("coping audio: %w", err)
		}
		err = copier.Copy(&audio, &v.Metadata)
		if err != nil {
			return fmt.Errorf("coping audio metadata: %w", err)
		}
		voc.Aud = append(voc.Aud, *audio)
	}
	return nil
}

func toVocKanji(revKanji *reviewKanji) (vacabularyKanji, error) {
	var vocKanji vacabularyKanji

	err := copier.Copy(&vocKanji, &revKanji)
	if err != nil {
		return vocKanji, fmt.Errorf("coping vocKanji: %w", err)
	}
	if revKanji.Type == common.ReadingTypeKunyomi {
		vocKanji.Ja = revKanji.Kun[0]
	} else if revKanji.Type == common.ReadingTypeNanori {
		vocKanji.Ja = revKanji.Nanori[0]
	} else if revKanji.Type == common.ReadingTypeOnyomi {
		vocKanji.Ja = revKanji.On[0]
	}
	return vocKanji, nil
}

func startVoc() reviewVocabulary {
	return reviewVocabulary{
		En:                []string{},
		Kana:              []string{},
		AuxiliaryMeanings: []common.AuxiliaryMeaning{},
		AuxiliaryReadings: []webapi.AuxiliaryReadings{},
		Aud:               []aud{},
		Kanji:             []vacabularyKanji{},
		Syn:               make([]interface{}, 0),
	}
}
