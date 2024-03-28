package json

import (
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/vocabulary"

	"github.com/jinzhu/copier"
	"github.com/samonzeweb/godb"
)

type jsonVocabulary struct {
	Type               string                           `json:"type"`
	ID                 int                              `json:"id"`
	Stroke             int                              `json:"stroke"`
	MeaningExplanation string                           `json:"meaning_explanation" copier:"MeaningMnemonic"`
	ReadingExplanation string                           `json:"reading_explanation" copier:"ReadingMnemonic"`
	Voc                string                           `json:"voc"`
	Characters         string                           `json:"characters"`
	En                 string                           `json:"en"`
	Kana               string                           `json:"kana"`
	Sentences          [][]string                       `json:"sentences"`
	MeaningNote        interface{}                      `json:"meaning_note"` //TODO: string?
	ReadingNote        interface{}                      `json:"reading_note"` //TODO: string?
	PartsOfSpeech      []string                         `json:"parts_of_speech"`
	Audio              []vocabulary.PronunciationAudios `json:"audio"`
	Related            []Related                        `json:"related"`
}

type Related struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Kan        string `json:"kan"`
	Characters string `json:"characters"`
	En         string `json:"en"`
	Slug       string `json:"slug"`
}

func getVocabulary(db *godb.DB, id int) (*jsonVocabulary, error) {
	voc, err := vocabulary.SelectVocabulary(db, id)
	if err != nil {
		return nil, err
	}

	var json jsonVocabulary
	err = copier.Copy(&json, voc)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&json, &voc.Data)
	if err != nil {
		return nil, err
	}
	json.Voc = json.Characters
	json.Stroke = voc.Data.Level // Is this correct?

	json.En = ""
	for _, v := range voc.Data.Meanings {
		if v.AcceptedAnswer {
			if len(json.En) > 0 {
				json.En += ", "
			}
			json.En += v.Meaning
		}
	}

	for _, v := range voc.Data.Readings {
		if v.Primary {
			json.Kana = v.Reading
		}
	}

	json.Sentences = make([][]string, len(voc.Data.ContextSentences))
	for i, sentence := range voc.Data.ContextSentences {
		json.Sentences[i] = []string{sentence.Ja, sentence.En}
	}

	json.Audio = voc.Data.PronunciationAudios

	json.Related = make([]Related, len(voc.Data.ComponentSubjectIds))
	sKanji := &kanji.Json{}
	for i, componentID := range voc.Data.ComponentSubjectIds {
		kanji, err := getSimplifedKanji(db, componentID, sKanji)
		if err != nil {
			return nil, err
		}
		err = copier.Copy(&json.Related[i], &kanji)
		if err != nil {
			return nil, err
		}
		json.Related[i].Slug = sKanji.Data.Slug
	}
	return &json, nil
}
