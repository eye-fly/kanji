package radical

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
func (json *Json) GetComponentSubjectIds() []int    { return []int{} }

type Data struct {
	CreatedAt                time.Time                    `json:"created_at" db:"created_at"`
	Level                    int                          `json:"level" db:"level"`
	Slug                     string                       `json:"slug" db:"slug"`
	HiddenAt                 time.Time                    `json:"hidden_at" db:"hidden_at"`
	DocumentURL              string                       `json:"document_url" db:"document_url"`
	Characters               string                       `json:"characters" db:"characters"`
	CharacterImages          []CharacterImages            `json:"character_images"`
	Meanings                 []subjects.Meanings          `json:"meanings" `
	AuxiliaryMeanings        []subjects.AuxiliaryMeanings `json:"auxiliary_meanings" `
	AmalgamationSubjectIds   []int                        `json:"amalgamation_subject_ids"`
	MeaningMnemonic          string                       `json:"meaning_mnemonic" db:"meaning_mnemonic"`
	LessonPosition           int                          `json:"lesson_position" db:"lesson_position"`
	SpacedRepetitionSystemID int                          `json:"spaced_repetition_system_id" db:"spaced_repetition_system_id"`
}
type Metadata struct {
	InlineStyles bool   `json:"inline_styles" db:"inline_styles"`
	Color        string `json:"color" db:"color"`
	Dimensions   string `json:"dimensions" db:"dimensions"`
	StyleName    string `json:"style_name" db:"style_name"`
}
type CharacterImages struct {
	RadicalId   int      `db:"radical_id"`
	URL         string   `json:"url" db:"url"`
	Metadata    Metadata `json:"metadata" db:""`
	ContentType string   `json:"content_type" db:"content_type"`
}

func (*CharacterImages) TableName() string { return characterImagesTable }
