package lesson

import (
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

type ComponentKanji struct {
	En         string `json:"en"`
	Ja         string `json:"ja"`
	Kan        string `json:"kan"`
	Slug       string `json:"slug"`
	Type       string `json:"type"`
	Characters string `json:"characters"`
}
type AuxiliaryReadings struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Reading string `json:"reading"`
}

type reations struct {
	StudyMaterial interface{} `json:"study_material"`
}

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
		Relationships:     reations{StudyMaterial: make([]int, 0)},
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

	for _, v := range voc.Data.Meanings {
		if v.AcceptedAnswer {
			json.En = append(json.En, v.Meaning)
		} else {
			json.AuxiliaryMeanings = append(json.AuxiliaryMeanings, common.AuxiliaryMeaning{
				Meaning: v.Meaning,
				Type:    "blacklist",
			})
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

	json.Sentences = make([][]string, len(voc.Data.ContextSentences))
	for i, sentence := range voc.Data.ContextSentences {
		json.Sentences[i] = []string{sentence.Ja, sentence.En}
	}

	json.Aud = make([]review.Aud, 0, len(voc.Data.PronunciationAudios))
	for _, v := range voc.Data.PronunciationAudios {
		json.Aud = append(json.Aud, review.Aud{
			URL:           v.URL,
			ContentType:   v.ContentType,
			Pronunciation: v.Metadata.Pronunciation,
			VoiceActorID:  v.Metadata.VoiceActorID,
		})
	}

	// json.Kanji = make([]componentKanji, len(voc.Data.ComponentSubjectIds))
	// sKanji := &kanji.Json{}
	// for i, componentID := range voc.Data.ComponentSubjectIds {
	// 	kanji, err := getSimplifedKanji(db, componentID, sKanji)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	err = copier.Copy(&json.Related[i], &kanji)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	json.Related[i].Slug = sKanji.Data.Slug
	// }
	return &json, nil
}
