package users

import (
	"fmt"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/common"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	"sql_filler/subjects/vocabulary"
	"time"

	"github.com/samonzeweb/godb"
	log "github.com/sirupsen/logrus"
)

func CanLevelup(db *godb.DB, userID int) error {
	var user User
	err := db.Select(&user).Where(UserIdColumn+" = ?", userID).Do()
	if err != nil {
		return err
	}
	if user.Level >= 60 {
		return nil
	}

	kan := make([]subjectID, 0)
	err = db.SelectFrom(kanji.KanjiTable).Where("level = ?", user.Level).
		Columns("id").Do(&kan)
	if err != nil {
		return err
	}

	var passed float32 = 0
	var notPassed float32 = 0
	for _, v := range kan {
		b, err := isPassed(db, userID, v.Id)
		if err != nil {
			return err
		} else if b {
			passed++
		} else {
			notPassed++
		}
	}

	levelCompleted := passed / (passed + notPassed)
	if levelCompleted >= 0.9 {
		user.Level++

		err = db.Update(&user).Do()
		if err != nil {
			return err
		}
		log.Infof("Level up to %vlvl of user with id: %v", user.Level, userID)
	}

	return nil
}

func UnlockLockedSubject(db *godb.DB, userID int) error {
	var user User
	err := db.Select(&user).Where(UserIdColumn+" = ?", userID).Do()
	if err != nil {
		return err
	}

	subjectIDs, err := selectAllSubj(db, user.Level, "<=")
	if err != nil {
		return err
	}

	for _, id := range subjectIDs {
		isUnloc, err := isUnlocked(db, userID, id.Id)
		if err != nil {
			return err
		} else if !isUnloc {
			canBeUnloc, err := CanBeUnlocked(db, userID, id.Id)
			if err != nil {
				return err
			} else if canBeUnloc {
				log.Infof("Unlocking subject with id: %v for user: %v", id.Id, userID)
				err = assignment.NewAssignment(db, userID, id.Id)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

type subjectID struct {
	Id int `db:"id"`
}

func selectAllSubj(db *godb.DB, level int, levelCmp string) (subjectIDs []subjectID, err error) {
	columns := fmt.Sprintf("SELECT %s, %s from %s WHERE %s "+levelCmp+" %v",
		"id", "hidden_at", "%s", common.LevelRow, level)
	tableUnion := fmt.Sprintf("("+columns+" UNION "+columns+" UNION "+columns+")",
		kanji.KanjiTable, radical.RadicalTable, vocabulary.VocabularyTable)

	err = db.SelectFrom(tableUnion).
		Where("hidden_at = ?", time.Time{}).Do(&subjectIDs)

	return
}

func isPassed(db *godb.DB, userID, subjectID int) (bool, error) {
	var row assignment.Json
	err := db.Select(&row).Where(assignment.SubjectIdRow+" = ? AND "+assignment.UserIdRow+" = ?",
		subjectID, userID).Do()

	return assignment.IsAfterStage(row.Data.PassedAt, err)
}

func isUnlocked(db *godb.DB, userID, subjectID int) (bool, error) {
	var row assignment.Json
	err := db.Select(&row).Where(assignment.SubjectIdRow+" = ? AND "+assignment.UserIdRow+" = ?",
		subjectID, userID).Do()

	return assignment.IsAfterStage(row.Data.UnlockedAt, err)
}

func CanBeUnlocked(db *godb.DB, userID, subjectID int) (bool, error) {
	type amalgamationSubjectComponent struct {
		Id int `db:"component_id"`
	}

	amalgamationSubjectComponents := make([]amalgamationSubjectComponent, 0)
	err := db.SelectFrom(common.AmalgamationSubjectTable).
		Where(common.AmalgamationSubjectSetRow+" = ?", subjectID).
		Columns(common.AmalgamationSubjectComponentRow).Do(&amalgamationSubjectComponents)
	if err != nil {
		return false, err
	}

	var row assignment.Json
	isAllPassed := true
	for _, v := range amalgamationSubjectComponents {
		err := db.Select(&row).Where(assignment.SubjectIdRow+" = ? AND "+assignment.UserIdRow+" = ?",
			v.Id, userID).Do()

		b, err := assignment.IsAfterStage(row.Data.PassedAt, err)
		if err != nil || !b {
			return false, err
		}
	}
	return isAllPassed, nil
}
