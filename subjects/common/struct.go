package common

type resource interface {
	TableName() string
	GetId() int
	GetType() string
}

type subject interface {
	TableName() string
	GetId() int
	GetType() string
	GetMeanings() []Meanings
	GetAuxiliaryMeanings() []AuxiliaryMeaning
	GetAmalgamationSubjectIds() []int
	GetComponentSubjectIds() []int
}

type Meanings struct {
	SubjectId      int    `db:"sybject_id"`
	Meaning        string `json:"meaning" db:"meaning"`
	Primary        bool   `json:"primary" db:"is_primary"`
	AcceptedAnswer bool   `json:"accepted_answer" db:"accepted_answer"`
}

func (*Meanings) TableName() string { return MeaningTable }

type AuxiliaryMeaning struct {
	SubjectId int    `db:"sybject_id"`
	Meaning   string `json:"meaning" db:"meaning"`
	Type      string `json:"type" db:"type"`
}

func (*AuxiliaryMeaning) TableName() string { return AuxiliaryMeaningTable }

type AmalgamationSubjectIds struct {
	ComponentId int `db:"component_id"`
	SetId       int `db:"set_id"`
}

func (*AmalgamationSubjectIds) TableName() string { return AmalgamationSubjectTable }

func NewRelation(componentId, setId int) *AmalgamationSubjectIds {
	return &AmalgamationSubjectIds{
		ComponentId: componentId,
		SetId:       setId,
	}
}

type Readings struct {
	SubjectId      int    `db:"subject_id"`
	Type           string `json:"type" db:"type"`
	Primary        bool   `json:"primary" db:"is_primary"`
	AcceptedAnswer bool   `json:"accepted_answer" db:"accepted_answer"`
	Reading        string `json:"reading" db:"reading"`
}

func (*Readings) TableName() string { return ReadingTable }
