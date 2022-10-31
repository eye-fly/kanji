package lesson

import (
	"fmt"
	"sql_filler/subjects/common"
	"sql_filler/subjects/kanji"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type lessonKan struct {
	En                []string                  `json:"en"`
	ID                int                       `json:"id"`
	Kan               string                    `json:"voc"`
	Kana              []string                  `json:"kana"`
	MeaningMnemonic   string                    `json:"mmne"`
	ReadingMnemonic   string                    `json:"rmne"`
	Type              string                    `json:"type"`
	Vocabulary        []ComponentVocabulary     `json:"vocabulary"`
	Radical           []ComponentRadical        `json:"radical"`
	Characters        string                    `json:"characters"`
	AuxiliaryMeanings []common.AuxiliaryMeaning `json:"auxiliary_meanings"`
	AuxiliaryReadings []AuxiliaryReadings       `json:"auxiliary_readings"`
	Level             int                       `json:"-"`
	LessonPosition    int                       `json:"-"`
	Relationships     reations                  `json:"relationships"`
}

func (j *lessonKan) getLevel() int          { return j.Level }
func (j *lessonKan) getLessonPosition() int { return j.LessonPosition }

func getKanji(db *godb.DB, id int) (*lessonKan, error) {
	voc, err := kanji.SelctKanji(db, id)
	if err != nil {
		return nil, err
	}

	json := lessonKan{
		Kana:              make([]string, 0),
		En:                make([]string, 0),
		Vocabulary:        make([]ComponentVocabulary, 0),
		Radical:           make([]ComponentRadical, 0),
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
	json.Kan = json.Characters
	json.Type = "Kanji"

	for _, v := range voc.Data.Meanings {
		if v.AcceptedAnswer {
			json.En = append(json.En, v.Meaning)
		}
	}

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

	json.Radical = make([]ComponentRadical, 0)
	for _, componentID := range voc.Data.ComponentSubjectIds {
		rad, err := getComponentRadical(db, componentID)
		if err == nil {
			json.Radical = append(json.Radical, *rad)
		} else {
			fmt.Printf("error geting component Radical: %v as componet of: %v| err: %s", componentID, id, err)
		}
	}
	json.Vocabulary = make([]ComponentVocabulary, 0)
	for _, componentID := range voc.Data.AmalgamationSubjectIds {
		voc, err := getComponentVocabulary(db, componentID)
		if err == nil {
			json.Vocabulary = append(json.Vocabulary, *voc)
		} else {
			fmt.Printf("error geting component Radical: %v as componet of: %v| err: %s", componentID, id, err)
		}
	}
	return &json, nil
}
