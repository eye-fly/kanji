package json

import (
	"errors"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type jsonKanji struct {
	Type            string           `json:"type"`
	ID              int              `json:"id"`
	Kan             string           `json:"kan"`
	Characters      string           `json:"characters"`
	Stroke          int              `json:"stroke"`
	MeaningMnemonic string           `json:"meaning_mnemonic"`
	MeaningHint     string           `json:"meaning_hint"`
	ReadingMnemonic string           `json:"reading_mnemonic"`
	ReadingHint     string           `json:"reading_hint"`
	En              string           `json:"en"`
	MeaningNote     interface{}      `json:"meaning_note"`
	ReadingNote     interface{}      `json:"reading_note"`
	Related         []RelatedRadical `json:"related"`
}
type RelatedRadical struct {
	Type              string      `json:"type"`
	ID                int         `json:"id"`
	Rad               string      `json:"rad"`
	Characters        string      `json:"characters"`
	En                string      `json:"en"`
	Slug              string      `json:"slug"`
	CharacterImageURL interface{} `json:"character_image_url"`
}

func getKanji(db *godb.DB, id int) (*jsonKanji, error) {
	sKanji := &kanji.Json{}
	json, err := getSimplifedKanji(db, id, sKanji)
	if err != nil {
		return nil, err
	}

	json.Related = make([]RelatedRadical, len(sKanji.Data.ComponentSubjectIds))
	for i, componentID := range sKanji.Data.ComponentSubjectIds {
		radcal, err := getRadical(db, componentID)
		if err != nil {
			return nil, err
		}
		err = copier.Copy(&json.Related[i], &radcal)
		if err != nil {
			return nil, err
		}

		if json.Related[i].Characters == "" {
			img := make([]radical.CharacterImages, 0)
			err = db.Select(&img).Where(radical.CharacterImagesIdRow+" = ?", componentID).
				Do()
			if err != nil {
				return nil, err
			}
			if len(img) == 0 || img[0].URL == "" {
				return nil, errors.New("get radical: no character nor img")
			}
			json.Related[i].CharacterImageURL = img[0].URL
		}
	}
	return json, nil
}

func getSimplifedKanji(db *godb.DB, id int, sKanji *kanji.Json) (*jsonKanji, error) {
	kanji, err := kanji.SelctKanji(db, id)
	if err != nil {
		return nil, err
	}
	*sKanji = *kanji

	var json jsonKanji
	copier.Copy(&json, kanji)
	if err != nil {
		return nil, err
	}
	copier.Copy(&json, &kanji.Data)
	if err != nil {
		return nil, err
	}
	json.Kan = json.Characters
	json.Stroke = kanji.Data.Level // Is this correct?

	for _, v := range kanji.Data.Meanings {
		if v.Primary {
			json.En = v.Meaning
			break
		}
	}

	return &json, err
}
