package assignment

import (
	"database/sql"
	"errors"
	"fmt"
	"sql_filler/subjects"
	"time"

	"github.com/samonzeweb/godb"
)

func (ass *Json) AddToDB(db *godb.DB, replace ...string) error {
	pres, err := ass.checkIfPresent(db)
	if err != nil {
		return err
	}

	if !pres {
		err = db.Insert(ass).Do()
	} else {
		err = db.Update(ass).Do()
	}
	return err
}

func GetFromDB(db *godb.DB, userID, subjectID int) (ass Json, err error) {
	err = db.Select(&ass).Where(SubjectIdRow+" = ? AND "+UserIdRow+" = ?", subjectID, userID).Do()
	if errors.Is(err, sql.ErrNoRows) {
		return ass, fmt.Errorf("updated not existing assignment: %w", err)
	}
	return
}

func GetSrsStage(db *godb.DB, userID, subjectID int) (int, error) {
	ass, err := GetFromDB(db, userID, subjectID)
	return ass.Data.SrsStage, err
}

type AssignmentDStage bool

const AssignmentDUp = true
const AssignmentDDown = true

func UpdateAssignment(db *godb.DB, userID, subjectID int, delta AssignmentDStage) error {
	item, err := GetFromDB(db, userID, subjectID)
	if err != nil {
		return err
	} else if item.Data.UnlockedAt.After(time.Now()) {
		return fmt.Errorf("can't change assigment before it's Unloced")
	} else if item.Data.AvailableAt.After(time.Now()) {
		return fmt.Errorf("can't change assigment before it's Available")
	}
	ass := &item

	nextStage := ass.Data.SrsStage
	if delta {
		if nextStage == 9 {
			return fmt.Errorf("can't go over burn satge")
		}
		nextStage++
	} else {
		nextStage--
	}
	// updata srsStage
	ass.Data.SrsStage = nextStage

	///updata times of compleating stages
	ass.checkIfPassed(nextStage)

	if !ass.Data.BurnedAt.IsZero() {
		//update 'avaiable at'
		d, err := subjects.GetLockDuration(db, subjectID, nextStage)
		if err != nil {
			return err
		}
		ass.Data.AvailableAt = time.Now().Add(d)
		ass.DataUpdatedAt = time.Now()
	}

	return db.Update(ass).Do()
}

// creates new assigment subject that is already unlocked
// but didn't yes pass lesson stage
func NewAssignment(db *godb.DB, userId, subjectID int, delta AssignmentDStage) error {
	subjectType, err := subjects.GetType(db, subjectID)
	if err != nil {
		return err
	}

	ass := Json{
		UserId:        userId,
		Object:        ObjectAssignment,
		DataUpdatedAt: time.Now(),
		Data: Data{
			CreatedAt:   time.Now(),
			SubjectID:   subjectID,
			SubjectType: string(subjectType),
			SrsStage:    0, //Start pre_lesson
			AvailableAt: time.Now(),
			Hidden:      false,
		},
	}

	return ass.AddToDB(db, "n")
}

func (ass *Json) checkIfPassed(nextStage int) {
	if ass.Data.StartedAt.IsZero() && nextStage >= 1 {
		ass.Data.StartedAt = time.Now()
		ass.DataUpdatedAt = time.Now()
	}
	if ass.Data.PassedAt.IsZero() && nextStage >= 5 {
		ass.Data.PassedAt = time.Now()
		ass.DataUpdatedAt = time.Now()
	}
	if ass.Data.BurnedAt.IsZero() && nextStage >= 9 {
		ass.Data.PassedAt = time.Now()
		ass.DataUpdatedAt = time.Now()
	}
}

func (ass *Json) checkIfPresent(db *godb.DB) (bool, error) {
	type countByAuthor struct {
		Count int `db:"count"`
	}
	var cout countByAuthor

	err := db.SelectFrom(ass.TableName()).Columns("count(*) as count").
		Where(UserIdRow+" = ? AND "+SubjectIdRow+" = ?", ass.UserId, ass.Data.SubjectID).
		Do(&cout)
	if err != nil {
		return false, err
	} else if cout.Count > 0 {
		return true, nil
	}
	return false, nil
}
