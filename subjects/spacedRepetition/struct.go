package spaced_repetition

import (
	"fmt"
	"time"
)

type Json struct {
	ID            int       `json:"id" db:"id,key"`
	Object        string    `json:"object" db:"object"`
	URL           string    `json:"url" db:"url"`
	DataUpdatedAt time.Time `json:"data_updated_at" db:"data_updated_at"`
	Data          Data      `json:"data" db:""`
}
type Data struct {
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	Name                   string    `json:"name" db:"name"`
	Description            string    `json:"description" db:"description"`
	UnlockingStagePosition int       `json:"unlocking_stage_position" db:"unlocking_stage_position"`
	StartingStagePosition  int       `json:"starting_stage_position" db:"starting_stage_position"`
	PassingStagePosition   int       `json:"passing_stage_position" db:"passing_stage_position"`
	BurningStagePosition   int       `json:"burning_stage_position" db:"burning_stage_position"`
	Stages                 []Stage   `json:"stages" `
}

func (*Json) TableName() string    { return srsTable }
func (json *Json) GetId() int      { return json.ID }
func (json *Json) GetType() string { return json.Object }

type Stage struct {
	SrsId        int    `db:"srs_id"`
	Interval     int    `json:"interval" db:"interval"`
	Position     int    `json:"position" db:"position"`
	IntervalUnit string `json:"interval_unit" db:"interval_unit"`
}

func (*Stage) TableName() string { return stageTable }

type IntervalU string

const IntervalMiliseconds IntervalU = "miliseconds"
const IntervalSeconds IntervalU = "seconds"
const IntervalMinutes IntervalU = "minutes"
const IntervalHours IntervalU = "hours"
const IntervalDays IntervalU = "days"
const IntervalWeeks IntervalU = "weeks"

func (interval IntervalU) ToTime() (time.Duration, error) {
	if interval == IntervalMiliseconds {
		return time.Millisecond, nil
	} else if interval == IntervalSeconds {
		return time.Second, nil
	} else if interval == IntervalMinutes {
		return time.Minute, nil
	} else if interval == IntervalHours {
		return time.Hour, nil
	} else if interval == IntervalDays {
		return 24 * time.Hour, nil
	} else if interval == IntervalWeeks {
		return 30 * 24 * time.Hour, nil
	}
	return 0, fmt.Errorf("to time fail")
}
