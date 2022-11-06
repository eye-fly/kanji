package review

import (
	"fmt"
	"math/rand"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/common"
	"time"

	"github.com/samonzeweb/godb"
)

type id struct {
	Nr int `db:"subject_id"`
}

func GetQueue(db *godb.DB, user_id int) ([]int, error) {
	ids := make([]id, 0)
	err := db.SelectFrom(assignment.AssignmentTable).
		Columns(assignment.SubjectIdRow).
		Where(fmt.Sprintf("%s = ? AND %s < ? AND %s > ?", assignment.UserIdRow, assignment.StartedAtRow, assignment.StartedAtRow), user_id, time.Now(), time.Time{}).
		Having(fmt.Sprintf("%s < ? AND %s > ? ", assignment.AvailableAtRow, assignment.AvailableAtRow), time.Now(), time.Time{}).
		GroupBy(assignment.SubjectIdRow).
		OrderBy(assignment.AvailableAtRow + "," + assignment.SubjectIdRow).
		Do(&ids)
	if err != nil {
		return nil, err
	}

	vsm := make([]int, len(ids))
	for i, v := range ids {
		vsm[i] = v.Nr
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vsm), func(i, j int) { vsm[i], vsm[j] = vsm[j], vsm[i] })
	return vsm, nil
}

func getAuxData(db *godb.DB, id int) (en []string,
	auxMeaning []common.AuxiliaryMeaning, err error) {

	meanings := make([]common.Meanings, 0) //////////////////////// ^imortant
	err = db.Select(&meanings).Where(common.SubjectIdRow+" = ? AND accepted_answer = 1", id).
		OrderBy("is_primary DESC").Do()
	if err != nil {
		return
	}
	en = make([]string, len(meanings))
	for i, meaning := range meanings {
		en[i] = meaning.Meaning
	}

	auxMeaning = make([]common.AuxiliaryMeaning, 0)
	err = db.Select(&auxMeaning).Where(common.SubjectIdRow+" = ?", id).Do()
	return

}
