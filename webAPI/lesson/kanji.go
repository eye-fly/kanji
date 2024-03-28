package lesson

import (
	"fmt"
	"math"
	"sql_filler/subjects/common"
	"sql_filler/subjects/kanji"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
	log "github.com/sirupsen/logrus"
)

type lessonKan struct {
	En                []string                  `json:"en"`
	ID                int                       `json:"id"`
	Kan               string                    `json:"kan"`
	Emph              string                    `json:"emph"`
	Kun               []string                  `json:"kun"`
	On                []string                  `json:"on"`
	Nanori            []string                  `json:"nanori"`
	MeaningMnemonic   string                    `json:"mmne"`
	ReadingMnemonic   string                    `json:"rmne"`
	MeaningHint       string                    `json:"mhnt"`
	ReadingHint       string                    `json:"rhnt"`
	Type              string                    `json:"type"`
	Vocabulary        []ComponentVocabulary     `json:"vocabulary"`
	Radical           []ComponentRadical        `json:"radicals"`
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
		En:                make([]string, 0),
		Kun:               make([]string, 0),
		On:                make([]string, 0),
		Nanori:            make([]string, 0),
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
			ok := false
			if v.Type == common.ReadingTypeKunyomi {
				json.Kun = append(json.Kun, v.Reading)
				ok = true
			} else if v.Type == common.ReadingTypeOnyomi {
				json.On = append(json.On, v.Reading)
				ok = true
			} else if v.Type == common.ReadingTypeNanori {
				json.Nanori = append(json.Nanori, v.Reading)
				ok = true
			} else {
				log.Warn("wrong type for reading of kanji with id: %v and rading type:%s", id, v.Type)
			}
			if ok && v.Primary {
				json.Emph = v.Type
			}
		} else {
			json.AuxiliaryReadings = append(json.AuxiliaryReadings, AuxiliaryReadings{
				Type:    "blacklist",
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
	ln := int(math.Min(4, float64(len(voc.Data.AmalgamationSubjectIds))))
	for _, componentID := range voc.Data.AmalgamationSubjectIds[0:ln] {
		voc, err := getComponentVocabulary(db, componentID)
		if err == nil {
			json.Vocabulary = append(json.Vocabulary, *voc)
		} else {
			fmt.Printf("error geting component Radical: %v as componet of: %v| err: %s", componentID, id, err)
		}
	}
	return &json, nil
}
