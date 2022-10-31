package lesson

import (
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	"sql_filler/subjects/vocabulary"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type ComponentKanji struct {
	En         string `json:"en"`
	Ja         string `json:"ja"`
	Kan        string `json:"kan"`
	Slug       string `json:"slug"`
	Type       string `json:"type"`
	Characters string `json:"characters"`
}

type ComponentVocabulary struct {
	En         string `json:"en"`
	Ja         string `json:"ja"`
	Voc        string `json:"voc"`
	Slug       string `json:"slug"`
	Type       string `json:"type"`
	Characters string `json:"characters"`
}

type ComponentRadical struct {
	En                string `json:"en"`
	Rad               string `json:"rad"`
	Slug              string `json:"slug"`
	Type              string `json:"type"`
	Characters        string `json:"characters"`
	CharacterImageURL string `json:"character_image_url" copier:"-"`
}

type AuxiliaryReadings struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Reading string `json:"reading"`
}

type reations struct {
	StudyMaterial interface{} `json:"study_material"`
}

func getComponentKanji(db *godb.DB, id int) (*ComponentKanji, error) {
	item, err := kanji.SelctKanji(db, id)
	if err != nil {
		return nil, err
	}

	json := ComponentKanji{}
	err = copier.Copy(&json, item)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, item.Data)
	if err != nil {
		return nil, err
	}
	json.Kan = item.Data.Characters
	json.Type = "Kanji"

	for _, v := range item.Data.Meanings {
		if v.Primary {
			json.En = v.Meaning
		}
	}
	for _, v := range item.Data.Readings {
		if v.Primary {
			json.Ja = v.Reading
		}
	}

	return &json, nil
}

func getComponentVocabulary(db *godb.DB, id int) (*ComponentVocabulary, error) {
	item, err := vocabulary.SelectVocabulary(db, id)
	if err != nil {
		return nil, err
	}

	json := ComponentVocabulary{}
	err = copier.Copy(&json, item)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, item.Data)
	if err != nil {
		return nil, err
	}
	json.Voc = json.Characters

	for _, v := range item.Data.Meanings {
		if v.Primary {
			json.En = v.Meaning
		}
	}
	for _, v := range item.Data.Readings {
		if v.Primary {
			json.Ja = v.Reading
		}
	}

	return &json, nil
}

func getComponentRadical(db *godb.DB, id int) (*ComponentRadical, error) {
	item, err := radical.SelctRadical(db, id)
	if err != nil {
		return nil, err
	}

	json := ComponentRadical{}
	err = copier.Copy(&json, item)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, item.Data)
	if err != nil {
		return nil, err
	}
	json.Rad = json.Characters

	for _, v := range item.Data.Meanings {
		if v.Primary {
			json.En = v.Meaning
		}
	}

	// check
	if json.Characters == "" {
		for _, v := range item.Data.CharacterImages {
			if v.ContentType == "image/png" {
				json.CharacterImageURL = v.URL
				break
			}
		}
	}

	return &json, nil
}
