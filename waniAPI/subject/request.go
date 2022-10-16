package subject

import (
	"fmt"
	"net/http"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	"sql_filler/subjects/vocabulary"
	waniapi "sql_filler/waniAPI"
)

func RequestRadical(client *http.Client, parameters map[string]string) ([]radical.Json, error) {
	colection, err := requestColection(client, radicalType, &responseRadicalColection{}, parameters)
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
	colection, err := requestColection(client, kanjiType, &responseKanjiColection{}, parameters)
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

func RequestVocabulary(client *http.Client, parameters map[string]string) ([]vocabulary.Json, error) {
	colection, err := requestColection(client, vocabularyType, &responseVocabularyColection{}, parameters)
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

func requestColection(client *http.Client, subjectType string, colection responseColection,
	parameters map[string]string) ([]interface{}, error) {

	parameters["types"] = subjectType
	req, err := waniapi.BuildGetRequest(subjectResourceURL, parameters)
	if err != nil {
		return nil, fmt.Errorf("[RequestRadical] Build request fail %w", err)
	}

	err = waniapi.DoAndDecode(client, req, colection, radicalType)
	if err != nil {
		return nil, err
	}
	radicalColection := colection.getData()

	for colection.getNextPage() != "" {
		req, err = waniapi.RequestNextPage(colection.getNextPage())

		waniapi.DoAndDecode(client, req, &colection, radicalType)
		if err != nil {
			return nil, err
		}
		radicalColection = append(radicalColection, colection.getData()...)
	}

	return radicalColection, nil
}
