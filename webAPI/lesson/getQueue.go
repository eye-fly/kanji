package lesson

import (
	"fmt"
	"math/rand"
	"sort"
	"sql_filler/subjects"
	"sql_filler/subjects/assignment"
	"time"

	"github.com/samonzeweb/godb"
)

type subjectJson interface {
	getLevel() int
	getLessonPosition() int
}

type queueJson struct {
	Count struct {
		Rad int `json:"rad"`
		Kan int `json:"kan"`
		Voc int `json:"voc"`
	} `json:"count"`
	Queue []subjectJson `json:"queue"`
}

type queueCountJson struct {
	Count int `json:"count"`
}

func getLessonQueue(db *godb.DB, user_id int) (*queueJson, error) {
	ids, err := getLessonQueueId(db, user_id)
	if err != nil {
		return nil, err
	}

	qJson := queueJson{}
	qJson.Queue = make([]subjectJson, 0)
	for _, id := range ids {
		typ, err := subjects.GetType(db, id)
		if err != nil {
			return nil, err
		}

		var item subjectJson
		if typ == subjects.TypeRadical {
			item, err = getRadical(db, id)
			qJson.Count.Rad++
		} else if typ == subjects.TypeKanji {
			item, err = getKanji(db, id)
			qJson.Count.Kan++
		} else if typ == subjects.TypeVocabulary {
			item, err = getVocabulary(db, id)
			qJson.Count.Voc++
		} else {
			err = fmt.Errorf("no maching type")
		}

		if err != nil {
			return nil, err
		} else {
			qJson.Queue = append(qJson.Queue, item)
		}

	}

	sort.Slice(qJson.Queue, func(i, j int) bool {
		if qJson.Queue[i].getLevel() == qJson.Queue[j].getLevel() {
			return qJson.Queue[i].getLessonPosition() < qJson.Queue[j].getLessonPosition()
		}
		return qJson.Queue[i].getLevel() < qJson.Queue[j].getLevel()
	})

	return &qJson, nil
}

type id struct {
	Nr int `db:"subject_id"`
}

func getLessonQueueId(db *godb.DB, user_id int) ([]int, error) {
	ids := make([]id, 0)
	err := db.SelectFrom(assignment.AssignmentTable).
		Columns(assignment.SubjectIdRow).
		Where(fmt.Sprintf("%s = ?", assignment.UserIdRow), user_id).
		Having(fmt.Sprintf("%s = ? AND %s <= ? ", assignment.StartedAtRow, assignment.UnlockedAtRow), time.Time{}, time.Now()).
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
