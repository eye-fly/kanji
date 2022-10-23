package subjects

import (
	"fmt"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	spaced_repetition "sql_filler/subjects/spacedRepetition"
	"sql_filler/subjects/vocabulary"
	"time"

	"github.com/samonzeweb/godb"
)

type subType string

const TypeKanji subType = "kanji"
const TypeRadical subType = "radical"
const TypeVocabulary subType = "vocabulary"

func GetType(db *godb.DB, id int) (subType, error) {
	type objectType struct {
		Type string `db:"object"`
	}
	var oType objectType

	rows := fmt.Sprintf("SELECT %s, %s FROM %s", "id", "object", "%s")
	tableUnion := fmt.Sprintf("("+rows+" UNION "+rows+" UNION "+rows+")",
		kanji.KanjiTable, radical.RadicalTable, vocabulary.VocabularyTable)
	err := db.SelectFrom(tableUnion).Columns("object").Where("id = ?", id).
		Do(&oType)
	if err != nil {
		return "", err
	}

	if oType.Type == string(TypeKanji) {
		return TypeKanji, nil
	} else if oType.Type == string(TypeRadical) {
		return TypeRadical, nil
	} else if oType.Type == string(TypeVocabulary) {
		return TypeVocabulary, nil
	}
	return "", fmt.Errorf("error while geting type of object id: %v", id)
}

func GetLockDuration(db *godb.DB, subjectID, srsStage int) (time.Duration, error) {
	type objectSrs struct {
		SrsId int `db:"spaced_repetition_system_id"`
	}
	var oSrs objectSrs

	rows := fmt.Sprintf("SELECT %s, %s FROM %s", "id", "spaced_repetition_system_id", "%s")
	tableUnion := fmt.Sprintf("("+rows+" UNION "+rows+" UNION "+rows+")",
		kanji.KanjiTable, radical.RadicalTable, vocabulary.VocabularyTable)
	err := db.SelectFrom(tableUnion).Columns("spaced_repetition_system_id").Where("id = ?", subjectID).
		Do(&oSrs)
	if err != nil {
		return 0, err
	}

	var stage spaced_repetition.Stage
	err = db.Select(&stage).
		Where(spaced_repetition.SrsSrageIdRow+" = ? AND position = ?",
			oSrs.SrsId, srsStage).Do()
	if err != nil {
		return 0, err
	}

	tUnit := spaced_repetition.IntervalU(stage.IntervalUnit)
	t, err := tUnit.ToTime()
	if err != nil {
		return 0, err
	}
	return t * time.Duration(stage.Interval), nil
}
