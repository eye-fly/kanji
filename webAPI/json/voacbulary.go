package json

import "github.com/samonzeweb/godb"

type jsonVocabulary struct {
	Type               string      `json:"type"`
	ID                 int         `json:"id"`
	Stroke             int         `json:"stroke"`
	MeaningExplanation string      `json:"meaning_explanation"`
	ReadingExplanation string      `json:"reading_explanation"`
	Voc                string      `json:"voc"`
	Characters         string      `json:"characters"`
	En                 string      `json:"en"`
	Kana               string      `json:"kana"`
	Sentences          [][]string  `json:"sentences"`
	MeaningNote        interface{} `json:"meaning_note"` //TODO: string?
	ReadingNote        interface{} `json:"reading_note"` //TODO: string?
	PartsOfSpeech      []string    `json:"parts_of_speech"`
	Audio              []audio     `json:"audio"`
	Related            []related   `json:"related"`
}

type audio struct {
	URL      string `json:"url"`
	Metadata struct {
		Gender           string `json:"gender"`
		SourceID         int    `json:"source_id"`
		Pronunciation    string `json:"pronunciation"`
		VoiceActorID     int    `json:"voice_actor_id"`
		VoiceActorName   string `json:"voice_actor_name"`
		VoiceDescription string `json:"voice_description"`
	} `json:"metadata"`
	ContentType string `json:"content_type"`
}

type related struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Kan        string `json:"kan"`
	Characters string `json:"characters"`
	En         string `json:"en"`
	Slug       string `json:"slug"`
}

func GetKanji(db *godb.DB, id int)
