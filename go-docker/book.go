package main

import (
	"fmt"
	"time"
	"github.com/lib/pq"
)

func getBook(bookID int) (Book, error) {
	//Retrieve
	res := Book{}

	var id int
	var name string
	var lastname string
	var faculty string
	var carer string
	var carne int
	var publicationDate pq.NullTime

	err := db.QueryRow(`SELECT id, name, lastname, faculty, carer, carne, publication_date FROM users where id = $1`, bookID).Scan(&id, &name, &lastname, &faculty, &carer, &carne, &publicationDate)
	if err == nil {
		res = Book{ID: id, Name: name, LastName:lastname, Faculty:faculty, Carer:carer, Carne: carne, PublicationDate: publicationDate.Time}
	}

	return res, err
}

func allBooks() ([]Book, error) {
	//Retrieve
	users := []Book{}

	rows, err := db.Query(`SELECT id, name, lastname, faculty, carer, carne, publication_date FROM users order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var lastname string
			var faculty string
			var carer string
			var carne int
			var publicationDate pq.NullTime

			err = rows.Scan(&id, &name, &lastname, &faculty, &carer, &carne, &publicationDate)
			if err == nil {
				currentBook := Book{ID: id, Name: name, LastName:lastname, Faculty:faculty, Carer:carer, Carne: carne}
				if publicationDate.Valid {
					currentBook.PublicationDate = publicationDate.Time
				}

				users = append(users, currentBook)
			} else {
				return users, err
			}
		}
	} else {
		return users, err
	}

	return users, err
}

func insertBook(name, lastname string, faculty string, carer string, carne int, publicationDate time.Time) (int, error) {
	//Create
	var bookID int
	err := db.QueryRow(`INSERT INTO users(name, lastname, faculty, carer, carne, publication_date) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, name, lastname, faculty, carer, carne, publicationDate).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID, err
}

func updateBook(id int, name, lastname string, faculty string, carer string, carne int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE users set name=$1, lastname=$2, faculty=$3, carer=$4, carne=$5, publication_date=$6 where id=$7 RETURNING id`, name, lastname, faculty, carer, carne, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from users where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
