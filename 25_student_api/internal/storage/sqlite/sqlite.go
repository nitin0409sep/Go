package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // As we are not using it anywhere, it just for initialization so we put _
	"github.com/nitin0409sep/students-api/internal/config"
	"github.com/nitin0409sep/students-api/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			age INTERGER
		)
	
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO STUDENTS (name, email, age) VALUES(?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return lastId, nil
}


func (s *Sqlite) GetById(id int64) (types.Student , error){
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM STUDENTS WHERE id = ? LIMIT 1")

	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)


	if err != nil {

		if(err == sql.ErrNoRows) {
			return types.Student{}, fmt.Errorf("No Student Found with id %s", fmt.Sprint(id))
		}

		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}