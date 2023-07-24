package subject

import (
	"io"
	"net/http"
	"os"
	"sql_filler/subjects/vocabulary"
	"strings"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetAudioUrlAndDownloadIt(levels string) {
	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	panicIfErr(err, db)

	client := &http.Client{}
	colection, err := RequestVocabulary(client, map[string]string{
		"levels": levels,
	})
	panicIfErr(err, db)

	for j, subjct := range colection {
		dbVoc, err := vocabulary.SelectVocabulary(db, subjct.ID)
		if err != nil {
			println("skiping: ", subjct.GetMeanings(), " not in db, err=", err.Error())
		} else {
			dbVoc.Data.PronunciationAudios = subjct.Data.PronunciationAudios

			for i, audio := range dbVoc.Data.PronunciationAudios {
				//  audio.URL
				var name = strings.TrimLeft(audio.URL, "https://files.wanikani.com/")
				dbVoc.Data.PronunciationAudios[i].URL = "http://" + staticFilesHostAdress + "/" + Static_file_dir_name + "/" + name
				err = DownloadFile(Static_file_dir_name+"/"+name, audio.URL)

				println(j+1, "/", len(colection), " ", i+1, "/", len(dbVoc.Data.PronunciationAudios))
				if err != nil {
					println("err while downloading webm file=", err.Error(), " vocabulary id: ", audio.VocabularyId)
				}

			}

			// for _, audio := range subjct.Data.PronunciationAudios {
			// 	println(audio.URL)
			// }
			err = dbVoc.AddToDB(db, "y")
			panicIfErr(err, db)
		}

	}
}

func panicIfErr(err error, db *godb.DB) {
	if err != nil {
		db.Close()
	}
}
