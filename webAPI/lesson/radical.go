package lesson

import (
	"fmt"
	"math"
	"sql_filler/subjects/common"
	"sql_filler/subjects/radical"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type lessonRad struct {
	En                []string                  `json:"en"`
	ID                int                       `json:"id"`
	Rad               string                    `json:"rad"`
	MeaningMnemonic   string                    `json:"mmne"`
	Type              string                    `json:"type"`
	Kanji             []ComponentKanji          `json:"kanji"`
	Characters        string                    `json:"characters"`
	AuxiliaryMeanings []common.AuxiliaryMeaning `json:"auxiliary_meanings"`
	CharacterImageURL string                    `json:"character_image_url"`
	Level             int                       `json:"-"`
	LessonPosition    int                       `json:"-"`
	Relationships     reations                  `json:"relationships"`
}

func (j *lessonRad) getLevel() int          { return j.Level }
func (j *lessonRad) getLessonPosition() int { return j.LessonPosition }

func getRadical(db *godb.DB, id int) (*lessonRad, error) {
	rad, err := radical.SelctRadical(db, id)
	if err != nil {
		return nil, err
	}

	json := lessonRad{
		En:                make([]string, 0),
		Kanji:             make([]ComponentKanji, 0),
		AuxiliaryMeanings: make([]common.AuxiliaryMeaning, 0),
		Relationships:     reations{},
	}
	err = copier.Copy(&json, rad)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, &rad.Data)
	if err != nil {
		return nil, err
	}
	json.Rad = json.Characters
	json.Type = "Radical"

	for _, v := range rad.Data.Meanings {
		if v.AcceptedAnswer {
			json.En = append(json.En, v.Meaning)
		}
	}

	crr_res := ""
	for _, v := range rad.Data.CharacterImages {
		if v.ContentType == "image/png" {
			if len(crr_res) < len(v.Metadata.Dimensions) {
				json.CharacterImageURL = v.URL
				crr_res = v.Metadata.Dimensions
			}
		}
	}

	json.Kanji = make([]ComponentKanji, 0)
	ln := int(math.Min(3, float64(len(rad.Data.AmalgamationSubjectIds))))
	for _, componentID := range rad.Data.AmalgamationSubjectIds[0:ln] {
		kan, err := getComponentKanji(db, componentID)
		if err == nil {
			json.Kanji = append(json.Kanji, *kan)
		} else {
			fmt.Printf("error geting component Kanji: %v as componet of: %v| err: %s", componentID, id, err)
		}
	}
	return &json, nil
}
