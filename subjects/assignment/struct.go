package assignment

import "time"

type Json struct {
	ID            int       `json:"id" db:"id,key"`
	UserId        int       `db:"user_id,key"`
	Object        string    `json:"object" db:"object"`
	URL           string    `json:"url" db:"url"`
	DataUpdatedAt time.Time `json:"data_updated_at" db:"data_updated_at"`
	Data          Data      `json:"data" db:""`
}
type Data struct {
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	SubjectID     int       `json:"subject_id" db:"subject_id"`
	SubjectType   string    `json:"subject_type" db:"subject_type"`
	SrsStage      int       `json:"srs_stage" db:"srs_stage"`
	UnlockedAt    time.Time `json:"unlocked_at" db:"unlocked_at"`
	StartedAt     time.Time `json:"started_at" db:"started_at"`
	PassedAt      time.Time `json:"passed_at" db:"passed_at"`
	BurnedAt      time.Time `json:"burned_at" db:"burned_at"`
	AvailableAt   time.Time `json:"available_at" db:"available_at"`
	ResurrectedAt time.Time `json:"resurrected_at" db:"resurrected_at"`
	Hidden        bool      `json:"hidden" db:"hidden"`
}

func (*Json) TableName() string    { return jsonTable }
func (json *Json) GetId() int      { return json.ID }
func (json *Json) GetType() string { return json.Object }
