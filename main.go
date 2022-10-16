package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sql_filler/subjects/radical"
	"sql_filler/waniAPI/subject"

	"log"
	"os"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

func tryUnmarshal() (*radical.Json, error) {
	myFile, err := os.Open("tst.json")
	if err != nil {
		return nil, err
	}
	defer myFile.Close()
	data, err := io.ReadAll(myFile)
	if err != nil {
		return nil, err
	}
	var result radical.Json
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {

	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	panicIfErr(err, db)
	// OPTIONAL: Set logger to show SQL execution logs
	db.SetLogger(log.New(os.Stderr, "", 0))

	// rdcal, err := tryUnmarshal()
	// panicIfErr(err, db)

	// err = rdcal.AddToDB(db)
	// panicIfErr(err, db)
	client := &http.Client{}
	colection, err := subject.RequestVocabulary(client, map[string]string{
		"levels": "1",
	})
	panicIfErr(err, db)
	for _, subjct := range colection {
		err = subjct.AddToDB(db, "y")
		panicIfErr(err, db)
	}

	/*
		// Multiple insert
		// Warning : BulkInsert only update ids with PostgreSQL and SQL Server!
		err = db.BulkInsert(&setTheLordOfTheRing).Do()
		panicIfErr(err)

		// Count
		count, err := db.SelectFrom("books").Count()
		panicIfErr(err)
		fmt.Println("Books count : ", count)

		// Custom select
		countByAuthor := make([]CountByAuthor, 0, 10)
		err = db.SelectFrom("books").
			Columns("author", "count(*) as count").
			GroupBy("author").
			Having("count(*) > 3").
			Do(&countByAuthor)
		fmt.Println("Count by authors : ", countByAuthor)

		// Select single object
		singleBook := Book{}
		err = db.Select(&singleBook).
			Where("title = ?", bookTheHobbit.Title).
			Do()
		if err == sql.ErrNoRows {
			// sql.ErrNoRows is only returned when the target is a single instance
			fmt.Println("Book not found !")
		} else {
			panicIfErr(err)
		}

		// Select single record values
		authorName := ""
		title := ""
		err = db.SelectFrom("books").
			Where("title = ?", bookTheHobbit.Title).
			Columns("author", "title").
			Scanx(&authorName, &title)
		if err == sql.ErrNoRows {
			// sql.ErrNoRows is only returned when the target is a single instance
			fmt.Println("Book not found !")
		} else {
			panicIfErr(err)
		}

		// Select multiple objects
		multipleBooks := make([]Book, 0, 0)
		err = db.Select(&multipleBooks).Do()
		panicIfErr(err)
		fmt.Println("Books found : ", len(multipleBooks))

		// Iterator
		iter, err := db.SelectFrom("books").
			Columns("id", "title", "author", "published").
			DoWithIterator()
		panicIfErr(err)
		for iter.Next() {
			book := Book{}
			err := iter.Scan(&book)
			panicIfErr(err)
			fmt.Println(book)
		}
		panicIfErr(iter.Err())
		panicIfErr(iter.Close())

		// Raw query
		subQuery := godb.NewSQLBuffer(0, 0). // sizes are indicative
							Write("select author ").
							Write("from books ").
							WriteCondition(godb.Q("where title = ?", bookTheHobbit.Title))

		queryBuffer := godb.NewSQLBuffer(64, 0).
			Write("select * ").
			Write("from books ").
			Write("where author in (").
			Append(subQuery).
			Write(")")

		panicIfErr(queryBuffer.Err())

		books := make([]Book, 0, 0)
		err = db.RawSQL(queryBuffer.SQL(), queryBuffer.Arguments()...).Do(&books)
		panicIfErr(err)
		fmt.Printf("Raw query found %d books\n", len(books))

		// Update and transactions
		err = db.Begin()
		panicIfErr(err)

		updated, err := db.UpdateTable("books").Set("author", "Tolkien").Do()
		panicIfErr(err)
		fmt.Println("Books updated : ", updated)

		bookTheHobbit.Author = "Tolkien"
		err = db.Update(&bookTheHobbit).Do()
		panicIfErr(err)
		fmt.Println("Books updated : ", updated)

		err = db.Rollback()
		panicIfErr(err)

		// Delete
		deleted, err := db.Delete(&bookTheHobbit).Do()
		panicIfErr(err)
		fmt.Println("Books deleted : ", deleted)

		deleted, err = db.DeleteFrom("books").
			WhereQ(godb.Or(
				godb.Q("author = ?", authorTolkien),
				godb.Q("author = ?", "Georged Orwell"),
			)).
			Do()
		panicIfErr(err)
		fmt.Println("Books deleted : ", deleted)
	*/
	// Bye
	err = db.Close()
	panicIfErr(err, nil)
}

// It's just an example, what did you expect ? (never do that in real code)
func panicIfErr(err error, db *godb.DB) {
	if err != nil {
		db.Close()
		panic(err)
	}
}
