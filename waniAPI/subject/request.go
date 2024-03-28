package subject

import (
	"fmt"
	"net/http"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	spaced_repetition "sql_filler/subjects/spacedRepetition"
	"sql_filler/subjects/vocabulary"
	waniapi "sql_filler/waniAPI"

	"github.com/samonzeweb/godb"
)

func GetAndPutAllSubjects(db *godb.DB, levelsString string) error {
	client := &http.Client{}

	mp := map[string]string{
		"levels": levelsString,
	}
	colection, err := RequestRadical(client, mp)
	if err != nil {
		return err
	}
	for _, subjct := range colection {
		err = subjct.AddToDB(db, "y")
		if err != nil {
			return err
		}
	}

	colectionK, err := RequestKanji(client, mp)
	if err != nil {
		return err
	}
	for _, subjct := range colectionK {
		err = subjct.AddToDB(db, "y")
		if err != nil {
			return err
		}
	}

	colectionV, err := RequestVocabulary(client, mp)
	if err != nil {
		return err
	}
	for _, subjct := range colectionV {
		err = subjct.AddToDB(db, "y")
		if err != nil {
			return err
		}
	}
	return nil
}

func RequestAssigment(client *http.Client, parameters map[string]string) ([]assignment.Json, error) {
	colection, err := requestArray(client, assaignmentResourceURL, &responseAssigmentColection{}, parameters)
	if err != nil {
		return nil, err
	}

	assignmentColectio := make([]assignment.Json, len(colection))
	var ok bool
	for i, v := range colection {
		assignmentColectio[i], ok = v.(assignment.Json)
		if !ok {
			return nil, fmt.Errorf("messed up casting type")
		}
	}
	return assignmentColectio, nil
}

func RequestRadical(client *http.Client, parameters map[string]string) ([]radical.Json, error) {
	parameters["types"] = radicalType
	colection, err := requestArray(client, subjectResourceURL, &responseRadicalColection{}, parameters)
	if err != nil {
		return nil, err
	}

	radicalColectio := make([]radical.Json, len(colection))
	var ok bool
	for i, v := range colection {
		radicalColectio[i], ok = v.(radical.Json)
		if !ok {
			return nil, fmt.Errorf("messed up casting type")
		}
	}
	return radicalColectio, nil
}
func RequestKanji(client *http.Client, parameters map[string]string) ([]kanji.Json, error) {
	parameters["types"] = kanjiType
	colection, err := requestArray(client, subjectResourceURL, &responseKanjiColection{}, parameters)
	if err != nil {
		return nil, err
	}

	kanjiColectio := make([]kanji.Json, len(colection))
	var ok bool
	for i, v := range colection {
		kanjiColectio[i], ok = v.(kanji.Json)
		if !ok {
			return nil, fmt.Errorf("messed up casting type")
		}
	}
	return kanjiColectio, nil
}

func RequestSpacedRepetition(client *http.Client, parameters map[string]string) ([]spaced_repetition.Json, error) {
	colection, err := requestArray(client, spacedRepetitionURL, &responseSpacedRepetitionColection{}, parameters)
	if err != nil {
		return nil, err
	}

	spacedRepetitionColectio := make([]spaced_repetition.Json, len(colection))
	var ok bool
	for i, v := range colection {
		spacedRepetitionColectio[i], ok = v.(spaced_repetition.Json)
		if !ok {
			return nil, fmt.Errorf("messed up casting type")
		}
	}
	return spacedRepetitionColectio, nil
}

func RequestVocabulary(client *http.Client, parameters map[string]string) ([]vocabulary.Json, error) {
	parameters["types"] = vocabularyType
	colection, err := requestArray(client, subjectResourceURL, &responseVocabularyColection{}, parameters)
	if err != nil {
		return nil, err
	}

	vocabularyColection := make([]vocabulary.Json, len(colection))
	var ok bool
	for i, v := range colection {
		vocabularyColection[i], ok = v.(vocabulary.Json)
		if !ok {
			return nil, fmt.Errorf("messed up casting type")
		}
	}
	return vocabularyColection, nil
}

func requestArray(client *http.Client, resourceURL string, colection responseWArray,
	parameters map[string]string) ([]interface{}, error) {

	req, err := waniapi.BuildGetRequest(resourceURL, parameters)
	if err != nil {
		return nil, fmt.Errorf("[RequestRadical] Build request fail %w", err)
	}

	println("url = ", req.URL.String())
	err = waniapi.DoAndDecode(client, req, colection, radicalType)
	if err != nil {
		return nil, err
	}
	println("getdatab")
	radicalColection := colection.getData()
	println("getdata af")

	for colection.getNextPage() != "" {
		println(colection.getNextPage())
		req, err = waniapi.RequestNextPage(colection.getNextPage())

		waniapi.DoAndDecode(client, req, &colection, radicalType)
		if err != nil {
			return nil, err
		}
		radicalColection = append(radicalColection, colection.getData()...)
	}

	return radicalColection, nil
}
