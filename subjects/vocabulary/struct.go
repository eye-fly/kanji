package vocabulary

import (
	"sql_filler/subjects"
	"time"
)

type Json struct {
	ID            int       `json:"id" db:"id,key"`
	Object        string    `json:"object" db:"object"`
	URL           string    `json:"url"  db:"url"`
	DataUpdatedAt time.Time `json:"data_updated_at" db:"data_updated_at"`
	Data          Data      `json:"data" db:""`
}

func (*Json) TableName() string                     { return jsonTable }
func (json *Json) GetId() int                       { return json.ID }
func (json *Json) GetType() string                  { return json.Object }
func (json *Json) GetMeanings() []subjects.Meanings { return json.Data.Meanings }
func (json *Json) GetAuxiliaryMeanings() []subjects.AuxiliaryMeanings {
	return json.Data.AuxiliaryMeanings
}
func (json *Json) GetAmalgamationSubjectIds() []int { return []int{} }
func (json *Json) GetComponentSubjectIds() []int    { return json.Data.ComponentSubjectIds }

type Data struct {
	AuxiliaryMeanings        []subjects.AuxiliaryMeanings `json:"auxiliary_meanings"`
	Characters               string                       `json:"characters" db:"characters"`
	ComponentSubjectIds      []int                        `json:"component_subject_ids"`
	ContextSentences         []ContextSentences           `json:"context_sentences"`
	CreatedAt                time.Time                    `json:"created_at" db:"created_at"`
	DocumentURL              string                       `json:"document_url" db:"document_url"`
	HiddenAt                 time.Time                    `json:"hidden_at" db:"hidden_at"`
	LessonPosition           int                          `json:"lesson_position" db:"lesson_position"`
	Level                    int                          `json:"level" db:"level"`
	Meanings                 []subjects.Meanings          `json:"meanings"`
	MeaningMnemonic          string                       `json:"meaning_mnemonic" db:"meaning_mnemonic"`
	PartsOfSpeech            []string                     `json:"parts_of_speech"`
	PronunciationAudios      []PronunciationAudios        `json:"pronunciation_audios"`
	Readings                 []subjects.Readings          `json:"readings"`
	ReadingMnemonic          string                       `json:"reading_mnemonic" db:"reading_mnemonic"`
	Slug                     string                       `json:"slug" db:"slug"`
	SpacedRepetitionSystemID int                          `json:"spaced_repetition_system_id" db:"spaced_repetition_system_id"`
}

type ContextSentences struct {
	En string `json:"en" db:"en"`
	Ja string `json:"ja" db:"ja"`
}
type Metadata struct {
	Gender           string `json:"gender" db:"gender"`
	SourceID         int    `json:"source_id" db:"source_id"`
	Pronunciation    string `json:"pronunciation" db:"pronunciation"`
	VoiceActorID     int    `json:"voice_actor_id" db:"voice_actor_id"`
	VoiceActorName   string `json:"voice_actor_name" db:"voice_actor_name"`
	VoiceDescription string `json:"voice_description" db:"voice_description"`
}
type PronunciationAudios struct {
	URL         string   `json:"url" db:"url"`
	Metadata    Metadata `json:"metadata" db:"metadata"`
	ContentType string   `json:"content_type" db:"content_type"`
}
