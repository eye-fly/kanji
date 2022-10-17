package subjects

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samonzeweb/godb"
)

type ReturnStatus string

const AddStatus ReturnStatus = "add"
const ReplaceStatus ReturnStatus = "replace"
const SkipStatus ReturnStatus = "skip"

func AddResourceDB(db *godb.DB, rsc resource, replace ...string) (ReturnStatus, error) {
	present, err := checkIfPresent(db, rsc)
	if err != nil {
		return SkipStatus, err
	}

	if !present {
		err = db.Insert(rsc).Do()
		if err != nil {
			return SkipStatus, err
		}
		return AddStatus, nil
	} else {
		if len(replace) == 0 {
			fmt.Printf("should %s id:%d be replaced? (%s - for Yes): ", rsc.GetType(), rsc.GetId(), ReplaceIfExist)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			replace = []string{input.Text()}
		}
		if replace[0] == ReplaceIfExist {
			err = db.Update(rsc).Do()
			if err != nil {
				return SkipStatus, err
			}

			return ReplaceStatus, nil
		}
		return SkipStatus, nil
	}
}

func AddSubjectDB(db *godb.DB, sb subject, replace ...string) (ReturnStatus, error) {

	status, err := AddResourceDB(db, sb, replace...)
	if err != nil {
		return SkipStatus, err
	}

	if status == AddStatus {
		err := addAuxularyData(db, sb)
		if err != nil {
			return SkipStatus, err
		}
	} else if status == ReplaceStatus {
		err = deleteAuxularyData(db, sb)
		if err != nil {
			return SkipStatus, err
		}
		err = addAuxularyData(db, sb)
		if err != nil {
			return SkipStatus, err
		}
	}
	return status, nil
}

func checkIfPresent(db *godb.DB, rsc resource) (bool, error) {
	type countByAuthor struct {
		Count int `db:"count"`
	}
	var cout countByAuthor

	err := db.SelectFrom(rsc.TableName()).
		Columns("count(*) as count").
		Where("id = ?", rsc.GetId()).Do(&cout)
	if err != nil {
		return false, err
	} else if cout.Count > 0 {
		return true, nil
	}
	return false, nil
}

func addAuxularyData(db *godb.DB, sb subject) error {
	for _, meaning := range sb.GetMeanings() {
		meaning.SubjectId = sb.GetId()
		err := db.Insert(&meaning).Do()
		if err != nil {
			return err
		}
	}
	for _, auxMeaning := range sb.GetAuxiliaryMeanings() {
		auxMeaning.SubjectId = sb.GetId()
		err := db.Insert(&auxMeaning).Do()
		if err != nil {
			return err
		}
	}

	sqlQ := fmt.Sprintf("INSERT OR IGNORE INTO \"%s\" (\"%s\", \"%s\") VALUES (?, ?) ",
		AmalgamationSubjectTable, AmalgamationSubjectComponentRow, AmalgamationSubjectSetRow)
	for _, setId := range sb.GetAmalgamationSubjectIds() {
		err := db.RawSQL(sqlQ, sb.GetId(), setId).Do(&[]AmalgamationSubjectIds{})
		if err != nil {
			return err
		}
	}
	for _, setId := range sb.GetComponentSubjectIds() {
		err := db.RawSQL(sqlQ, setId, sb.GetId()).Do(&[]AmalgamationSubjectIds{})
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteAuxularyData(db *godb.DB, sb subject) error {
	_, err := db.DeleteFrom(MeaningTable).WhereQ(
		godb.Q(SubjectIdRow+" = ?", sb.GetId()),
	).Do()
	if err != nil {
		return err
	}

	_, err = db.DeleteFrom(AuxiliaryMeaningTable).WhereQ(
		godb.Q(SubjectIdRow+" = ?", sb.GetId()),
	).Do()
	if err != nil {
		return err
	}

	_, err = db.DeleteFrom(AmalgamationSubjectTable).WhereQ(
		godb.Q(AmalgamationSubjectComponentRow+" = ?", sb.GetId()),
	).Do()
	if err != nil {
		return err
	}
	_, err = db.DeleteFrom(AmalgamationSubjectTable).WhereQ(
		godb.Q(AmalgamationSubjectSetRow+" = ?", sb.GetId()),
	).Do()
	if err != nil {
		return err
	}
	return nil
}
