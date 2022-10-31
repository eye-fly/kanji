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

func GetFromDB(db *godb.DB, userID, subjectID int) (ass *Json, err error) {
	ass = &Json{}
	err = db.Select(ass).Where(SubjectIdRow+" = ? AND "+UserIdRow+" = ?", subjectID, userID).Do()
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
const AssignmentDDown = false

func UpdateAssignment(db *godb.DB, userID, subjectID int, delta AssignmentDStage) error {
	ass, err := GetFromDB(db, userID, subjectID)
	if err != nil {
		return err
	}

	err = ass.processProgress(db, delta)
	if err != nil {
		return err
	}

	return db.Update(ass).Do()
}

func StartAssignment(db *godb.DB, userID, subjectID int) error {
	ass, err := GetFromDB(db, userID, subjectID)
	if err != nil {
		return err
	}

	if ass.Data.UnlockedAt.IsZero() || ass.Data.UnlockedAt.After(time.Now()) {
		return fmt.Errorf("can't change assigment before it's Unloced")
	} else if !ass.Data.StartedAt.IsZero() {
		return fmt.Errorf("assigment already started")
	} else

	// updata srsStage
	if ass.Data.SrsStage != 0 {
		return fmt.Errorf("bad assigment stage")
	}
	ass.Data.SrsStage++

	///updata times of compleating stages
	ass.Data.StartedAt = time.Now()

	//update 'avaiable at'
	d, err := subjects.GetLockDuration(db, ass.Data.SubjectID, ass.Data.SrsStage)
	if err != nil {
		return err
	}
	ass.Data.AvailableAt = time.Now().Add(d)

	ass.DataUpdatedAt = time.Now()
	return db.Update(ass).Do()
}

// creates new assigment subject that is already unlocked
// but didn't pass lesson stage
func NewAssignment(db *godb.DB, userId, subjectID int) error {
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

func (ass *Json) processProgress(db *godb.DB, delta AssignmentDStage) error {
	if ass.Data.UnlockedAt.IsZero() || ass.Data.UnlockedAt.After(time.Now()) {
		return fmt.Errorf("can't change assigment before it's Unloced")
	} else if ass.Data.AvailableAt.After(time.Now()) {
		return fmt.Errorf("can't change assigment before it's Available")
	}
	nextStage := ass.Data.SrsStage
	if delta {
		if nextStage == 9 {
			return fmt.Errorf("can't go over burn satge")
		}
		nextStage++
	} else {
		if nextStage == 0 {
			return nil
		} else if nextStage > 1 {
			nextStage--
		}
	}
	// updata srsStage
	ass.Data.SrsStage = nextStage

	///updata times of compleating stages
	ass.checkIfPassed(nextStage)

	if ass.Data.BurnedAt.IsZero() {
		//update 'avaiable at'
		d, err := subjects.GetLockDuration(db, ass.Data.SubjectID, nextStage)
		if err != nil {
			return err
		}
		ass.Data.AvailableAt = time.Now().Add(d)
	} else {
		ass.Data.AvailableAt = time.Time{}
	}
	ass.DataUpdatedAt = time.Now()

	return nil
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
		ass.Data.BurnedAt = time.Now()
		ass.DataUpdatedAt = time.Now()
	}
}

func IsAfterStage(stagePass time.Time, err error) (bool, error) {
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	if stagePass.IsZero() || !stagePass.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}
