package lesson

import (
	"fmt"
	"sql_filler/subjects/common"
	"sql_filler/subjects/vocabulary"
	"sql_filler/webAPI/review"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type lessonVoc struct {
	En                []string                  `json:"en"`
	ID                int                       `json:"id"`
	Aud               []review.Aud              `json:"aud"`
	Voc               string                    `json:"voc"`
	Kana              []string                  `json:"kana"`
	MeaningMnemonic   string                    `json:"mmne"`
	ReadingMnemonic   string                    `json:"rmne"`
	Type              string                    `json:"type"`
	Kanji             []ComponentKanji          `json:"kanji"`
	Sentences         [][]string                `json:"sentences"`
	Characters        string                    `json:"characters"`
	PartsOfSpeech     []string                  `json:"parts_of_speech"`
	AuxiliaryMeanings []common.AuxiliaryMeaning `json:"auxiliary_meanings"`
	AuxiliaryReadings []AuxiliaryReadings       `json:"auxiliary_readings"`
	Level             int                       `json:"-"`
	LessonPosition    int                       `json:"-"`
	Relationships     reations                  `json:"relationships"`
	Collocations      []colection               `json:"collocations"`
}

func (j *lessonVoc) getLevel() int          { return j.Level }
func (j *lessonVoc) getLessonPosition() int { return j.LessonPosition }

type colection struct {
	English      string `json:"english"`
	Japanese     string `json:"japanese"`
	PatternOfUse string `json:"pattern_of_use"`
}

func getVocabulary(db *godb.DB, id int) (*lessonVoc, error) {
	voc, err := vocabulary.SelectVocabulary(db, id)
	if err != nil {
		return nil, err
	}

	json := lessonVoc{
		Collocations:      make([]colection, 0),
		Aud:               make([]review.Aud, 0),
		Kana:              make([]string, 0),
		En:                make([]string, 0),
		Kanji:             make([]ComponentKanji, 0),
		PartsOfSpeech:     make([]string, 0),
		AuxiliaryMeanings: make([]common.AuxiliaryMeaning, 0),
		AuxiliaryReadings: make([]AuxiliaryReadings, 0),
		Relationships:     reations{},
	}
	err = copier.Copy(&json, voc)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, &voc.Data)
	if err != nil {
		return nil, err
	}
	json.Voc = json.Characters
	json.Type = "Vocabulary"

	for _, v := range voc.Data.Meanings {
		if v.AcceptedAnswer {
			json.En = append(json.En, v.Meaning)
		}
	}

	// use aux reading ...
	for _, v := range voc.Data.Readings {
		if v.AcceptedAnswer {
			if v.Primary {
				json.Kana = append(json.Kana, v.Reading)

			} else {
				json.AuxiliaryReadings = append(json.AuxiliaryReadings, AuxiliaryReadings{
					Type:    "whitelist",
					Message: "Check page ...",
					Reading: v.Reading,
				})
			}
		} else {
			json.AuxiliaryReadings = append(json.AuxiliaryReadings, AuxiliaryReadings{
				Type:    "warn",
				Message: "Check page ...",
				Reading: v.Reading,
			})
		}
	}

	json.Sentences = make([][]string, len(voc.Data.ContextSentences))
	for i, sentence := range voc.Data.ContextSentences {
		json.Sentences[i] = []string{sentence.Ja, sentence.En}
	}

	json.Aud = make([]review.Aud, 0)
	for _, v := range voc.Data.PronunciationAudios {
		json.Aud = append(json.Aud, review.Aud{
			URL:           v.URL,
			ContentType:   v.ContentType,
			Pronunciation: v.Metadata.Pronunciation,
			VoiceActorID:  v.Metadata.VoiceActorID,
		})
	}

	json.Kanji = make([]ComponentKanji, 0)
	for _, componentID := range voc.Data.ComponentSubjectIds {
		kan, err := getComponentKanji(db, componentID)
		if err == nil {
			json.Kanji = append(json.Kanji, *kan)
		} else {
			fmt.Printf("error geting component Kanji: %v as componet of: %v| err: %s", componentID, id, err)
		}
	}
	return &json, nil
}
