package learining

import (
	"net/http"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/kanji"
	"sql_filler/subjects/vocabulary"
	"sql_filler/webAPI/json"
)

func (bec *backEnd) findVoc(userID int, data findSubmitData, w http.ResponseWriter) {

	var vocs []*vocabulary.Json

	if data.SearchWriting {
		voc, err := vocabulary.FindVocabulary(bec.db, data.SearchText)
		if err == nil {
			vocs = append(vocs, voc)
		}
	}
	if data.SearchReading {
		vocsp, err := vocabulary.FindVocabularyByReading(bec.db, data.SearchText)
		if err == nil {
			vocs = append(vocs, vocsp...)
		}

	}

	for _, v := range vocs {
		ass, err := assignment.GetFromDB(bec.db, userID, v.ID)
		if err != nil {
			v.Data.IsBurned = false
			v.Data.IsLearning = false
		} else {
			if ass.Data.BurnedAt.IsZero() {
				v.Data.IsBurned = false
			} else {
				v.Data.IsBurned = true
			}

			v.Data.IsLearning = ass.Data.IsLearning
		}
	}

	if len(vocs) > 0 {
		json.ServeBodyJson(w, vocs)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (bec *backEnd) findKanji(userID int, data findSubmitData, w http.ResponseWriter) {

	var kan []*kanji.Json

	if data.SearchWriting {
		k, err := kanji.FindKanji(bec.db, data.SearchText)
		if err == nil {
			kan = append(kan, k)
		}
	}
	// TODO
	// if data.SearchReading {
	// 	vocsp, err := vocabulary.FindVocabularyByReading(bec.db, data.SearchText)
	// 	if err == nil {
	// 		vocs = append(vocs, vocsp...)
	// 	}

	// }

	// for _, v := range vocs {
	// 	ass, err := assignment.GetFromDB(bec.db, userID, v.ID)
	// 	if err == nil {
	// 		v.Data.IsBurned = false
	// 		v.Data.IsLearning = false
	// 	} else {
	// 		if ass.Data.BurnedAt.IsZero() {
	// 			v.Data.IsBurned = false
	// 		} else {
	// 			v.Data.IsBurned = true
	// 		}

	// 		v.Data.IsLearning = ass.Data.IsLearning
	// 	}
	// }

	if len(kan) > 0 {
		json.ServeBodyJson(w, kan)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
