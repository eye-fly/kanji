package review

import (
	"errors"
	"fmt"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/common"
	"sql_filler/subjects/radical"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type reviewRadical struct {
	En                []string                  `json:"en" copier:"-"`
	ID                int                       `json:"id"`
	Rad               string                    `json:"rad"`
	Type              string                    `json:"type"`
	Characters        string                    `json:"characters" copier:"Characters"`
	AuxiliaryMeanings []common.AuxiliaryMeaning `json:"auxiliary_meanings" copier:"-"`
	CharacterImageURL string                    `json:"character_image_url" copier:"-"`
	Srs               int                       `json:"srs" copier:"Srs"`
	Syn               []interface{}             `json:"syn" copier:"-"`
}

// const NoImg  := errors.New("no character nor img")

func GetRadical(db *godb.DB, subjectID, userID int) (*reviewRadical, error) {
	var item radical.Json
	err := db.Select(&item).Where("id = ?", subjectID).Do()
	if err != nil {
		return nil, err
	}
	var reviewItem reviewRadical
	err = copier.Copy(&reviewItem, &item)
	if err != nil {
		return nil, fmt.Errorf("coping radical: %w", err)
	}
	err = copier.Copy(&reviewItem, &item.Data)
	if err != nil {
		return nil, fmt.Errorf("coping radical data: %w", err)
	}
	reviewItem.Rad = reviewItem.Characters
	reviewItem.Type = "Radical"

	reviewItem.Srs, err = assignment.GetSrsStage(db, userID, subjectID)
	if err != nil {
		return nil, err
	}
	reviewItem.En, reviewItem.AuxiliaryMeanings, err = getAuxData(db, subjectID)
	if err != nil {
		return nil, fmt.Errorf("auxData radical: %w", err)
	}

	if reviewItem.Characters == "" {
		img := make([]radical.CharacterImages, 0)
		err = db.Select(&img).Where(radical.CharacterImagesIdRow+" = ? AND content_type = \"image/png\"", subjectID).
			Do()
		if err != nil {
			return nil, err
		}
		if len(img) == 0 || img[0].URL == "" {
			return nil, errors.New("get radical: no character nor img")
		}
		reviewItem.CharacterImageURL = img[0].URL
	}

	return &reviewItem, nil
}
