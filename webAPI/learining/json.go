package learining

type findSubmitData struct {
	Type          string `json:"type"`
	SearchReading bool   `json:"searchReading"`
	SearchWriting bool   `json:"searchWriting"`
	SearchText    string `json:"searchText"`
}

const findSubmitTypeVocabulary string = "Vocabulary"
const findSubmitTypeKanji string = "Kanji"
