package sqlite

import (
	"go-cleanarchitecture/internal/app/service"
	"log"

	// sqlite is used as the backend
	_ "github.com/mattn/go-sqlite3"
)

func (r *Server) Create(p service.Person) (err error) {

	tx, err := r.dbConn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into person (id, name, age) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.UUID, p.Name, p.Age)
	if err != nil {
		log.Fatal(err)
	}
	// insertID, _ := res.LastInsertId()
	tx.Commit()

	return nil
}

func (r *Server) FindAll() []service.Person {

	rows, err := r.dbConn.Query("SELECT id, name, age FROM person")

	if err != nil {
		log.Fatal("Unable to query foo table:", err)
	}
	defer rows.Close()

	var ps []service.Person
	for rows.Next() {
		var p service.Person

		if err := rows.Scan(&p.UUID, &p.Name, &p.Age); err != nil {
			log.Printf("Unable to scan results: %w", err)
			continue
		}
		ps = append(ps, p)
	}
	return ps
}
