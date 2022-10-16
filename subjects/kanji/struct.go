package kanji

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
func (json *Json) GetAmalgamationSubjectIds() []int { return json.Data.AmalgamationSubjectIds }
func (json *Json) GetComponentSubjectIds() []int    { return json.Data.ComponentSubjectIds }

type Data struct {
	AmalgamationSubjectIds    []int                        `json:"amalgamation_subject_ids"`
	AuxiliaryMeanings         []subjects.AuxiliaryMeanings `json:"auxiliary_meanings"`
	Characters                string                       `json:"characters" db:"characters"`
	ComponentSubjectIds       []int                        `json:"component_subject_ids"`
	CreatedAt                 time.Time                    `json:"created_at" db:"created_at"`
	DocumentURL               string                       `json:"document_url" db:"document_url"`
	HiddenAt                  time.Time                    `json:"hidden_at" db:"hidden_at"`
	LessonPosition            int                          `json:"lesson_position" db:"lesson_position"`
	Level                     int                          `json:"level" db:"level"`
	Meanings                  []subjects.Meanings          `json:"meanings"`
	MeaningHint               string                       `json:"meaning_hint" db:"meaning_hint"`
	MeaningMnemonic           string                       `json:"meaning_mnemonic" db:"meaning_mnemonic"`
	Readings                  []subjects.Readings          `json:"readings"`
	ReadingMnemonic           string                       `json:"reading_mnemonic" db:"reading_mnemonic"`
	ReadingHint               string                       `json:"reading_hint" db:"reading_hint"`
	Slug                      string                       `json:"slug" db:"slug"`
	VisuallySimilarSubjectIds []int                        `json:"visually_similar_subject_ids"`
	SpacedRepetitionSystemID  int                          `json:"spaced_repetition_system_id" db:"spaced_repetition_system_id"`
}
