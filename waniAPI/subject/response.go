package subject

import (
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/radical"
	spaced_repetition "sql_filler/subjects/spacedRepetition"
	"sql_filler/subjects/vocabulary"
	"time"
)

type responseWArray interface {
	getNextPage() string
	getData() []interface{}
}

type responseAssigmentColection struct {
	Object        string            `json:"object"`
	URL           string            `json:"url"`
	Pages         pages             `json:"pages"`
	TotalCount    int               `json:"total_count"`
	DataUpdatedAt time.Time         `json:"data_updated_at"`
	Data          []assignment.Json `json:"data"`
}

func (resp *responseAssigmentColection) getNextPage() string { return resp.Pages.NextURL }
func (resp *responseAssigmentColection) getData() []interface{} {
	y := make([]interface{}, len(resp.Data))
	for i, v := range resp.Data {
		y[i] = v
	}
	return y
}

type responseRadicalColection struct {
	Object        string         `json:"object"`
	URL           string         `json:"url"`
	Pages         pages          `json:"pages"`
	TotalCount    int            `json:"total_count"`
	DataUpdatedAt time.Time      `json:"data_updated_at"`
	Data          []radical.Json `json:"data"`
}

func (resp *responseRadicalColection) getNextPage() string { return resp.Pages.NextURL }
func (resp *responseRadicalColection) getData() []interface{} {
	y := make([]interface{}, len(resp.Data))
	for i, v := range resp.Data {
		y[i] = v
	}
	return y
}

type responseKanjiColection struct {
	Object        string       `json:"object"`
	URL           string       `json:"url"`
	Pages         pages        `json:"pages"`
	TotalCount    int          `json:"total_count"`
	DataUpdatedAt time.Time    `json:"data_updated_at"`
	Data          []kanji.Json `json:"data"`
}

func (resp *responseKanjiColection) getNextPage() string { return resp.Pages.NextURL }
func (resp *responseKanjiColection) getData() []interface{} {
	y := make([]interface{}, len(resp.Data))
	for i, v := range resp.Data {
		y[i] = v
	}
	return y
}

type responseSpacedRepetitionColection struct {
	Object        string                   `json:"object"`
	URL           string                   `json:"url"`
	Pages         pages                    `json:"pages"`
	TotalCount    int                      `json:"total_count"`
	DataUpdatedAt time.Time                `json:"data_updated_at"`
	Data          []spaced_repetition.Json `json:"data"`
}

func (resp *responseSpacedRepetitionColection) getNextPage() string { return resp.Pages.NextURL }
func (resp *responseSpacedRepetitionColection) getData() []interface{} {
	y := make([]interface{}, len(resp.Data))
	for i, v := range resp.Data {
		y[i] = v
	}
	return y
}

type responseVocabularyColection struct {
	Object        string            `json:"object"`
	URL           string            `json:"url"`
	Pages         pages             `json:"pages"`
	TotalCount    int               `json:"total_count"`
	DataUpdatedAt time.Time         `json:"data_updated_at"`
	Data          []vocabulary.Json `json:"data"`
}

func (resp *responseVocabularyColection) getNextPage() string { return resp.Pages.NextURL }
func (resp *responseVocabularyColection) getData() []interface{} {
	y := make([]interface{}, len(resp.Data))
	for i, v := range resp.Data {
		y[i] = v
	}
	return y
}

type pages struct {
	PerPage     int         `json:"per_page"`
	NextURL     string      `json:"next_url"`
	PreviousURL interface{} `json:"previous_url"`
}
