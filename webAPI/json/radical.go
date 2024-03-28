package json

import (
	"sql_filler/subjects/radical"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type jsonRadical struct {
	Type        string      `json:"type"`
	ID          int         `json:"id"`
	Rad         string      `json:"rad"`
	Characters  string      `json:"characters"`
	Stroke      int         `json:"stroke"`
	Mnemonic    string      `json:"mnemonic"`
	MeaningNote interface{} `json:"meaning_note"`
	En          string      `json:"en"`
	Slug        string
}

func getRadical(db *godb.DB, id int) (*jsonRadical, error) {
	radical, err := radical.SelctRadical(db, id)
	if err != nil {
		return nil, err
	}

	var json jsonRadical
	copier.Copy(&json, radical)
	if err != nil {
		return nil, err
	}
	copier.Copy(&json, &radical.Data)
	if err != nil {
		return nil, err
	}
	json.Rad = json.Characters
	json.Stroke = radical.Data.Level // Is this correct?
	json.Mnemonic = radical.Data.MeaningMnemonic

	for _, v := range radical.Data.Meanings {
		if v.Primary {
			json.En = v.Meaning
			break
		}
	}
	return &json, nil
}
