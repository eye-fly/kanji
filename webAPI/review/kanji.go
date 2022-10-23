package review

import (
	"fmt"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/common"
	"sql_filler/subjects/kanji"
	webapi "sql_filler/webAPI"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type reviewKanji struct {
	ID                int                        `json:"id"`
	Type              string                     `json:"type" copier:"Type"`
	Kan               string                     `json:"kan"`
	Characters        string                     `json:"characters" copier:"Characters"`
	En                []string                   `json:"en"`
	Emph              string                     `json:"emph"`
	On                []string                   `json:"on"`
	Kun               []string                   `json:"kun"`
	Nanori            []string                   `json:"nanori"`
	AuxiliaryMeanings []common.AuxiliaryMeaning  `json:"auxiliary_meanings"`
	AuxiliaryReadings []webapi.AuxiliaryReadings `json:"auxiliary_readings"`
	Srs               int                        `json:"srs" copier:"Srs"`
	Syn               []interface{}              `json:"syn"`
}

func GetKanji(db *godb.DB, subjectID, userID int) (*reviewKanji, error) {
	var item kanji.Json
	err := db.Select(&item).Where("id = ?", subjectID).Do()
	if err != nil {
		return nil, err
	}

	reviewItem := startKanji()
	err = copier.Copy(&reviewItem, &item)
	if err != nil {
		return nil, fmt.Errorf("coping kanji: %w", err)
	}
	err = copier.Copy(&reviewItem, &item.Data)
	if err != nil {
		return nil, fmt.Errorf("coping kanji data: %w", err)
	}
	reviewItem.Kan = reviewItem.Characters
	reviewItem.Type = "Kanji"

	reviewItem.Srs, err = assignment.GetSrsStage(db, userID, subjectID)
	if err != nil {
		return nil, err
	}
	reviewItem.En, reviewItem.AuxiliaryMeanings, err = getAuxData(db, subjectID)
	if err != nil {
		return nil, fmt.Errorf("auxData kanji: %w", err)
	}

	readings := make([]common.Readings, 0)
	err = db.Select(&readings).Where(common.ReadingsIdRow+" = ?", subjectID).
		Do()
	if err != nil {
		return nil, err
	}
	for _, reading := range readings {
		if reading.AcceptedAnswer {
			if reading.Primary {
				reviewItem.Emph = reading.Type
			}
			if reading.Type == common.ReadingTypeKunyomi {
				reviewItem.Kun = append(reviewItem.Kun, reading.Reading)
			} else if reading.Type == common.ReadingTypeNanori {
				reviewItem.Nanori = append(reviewItem.Nanori, reading.Reading)
			} else if reading.Type == common.ReadingTypeOnyomi {
				reviewItem.On = append(reviewItem.On, reading.Reading)
			}
		}
	}
	return &reviewItem, nil
}

func startKanji() reviewKanji {
	return reviewKanji{
		En:                []string{},
		On:                []string{},
		Kun:               []string{},
		Nanori:            []string{},
		AuxiliaryMeanings: []common.AuxiliaryMeaning{},
		AuxiliaryReadings: []webapi.AuxiliaryReadings{},
		Syn:               make([]interface{}, 0),
	}
}
