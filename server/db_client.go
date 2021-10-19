package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

////////////////////////////////////////////////////////
func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/learning?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

////////////////////////////////////////////////////////
func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS todo (
        id INT AUTO_INCREMENT,
        task TEXT NOT NULL,
        status TEXT NOT NULL,
        PRIMARY KEY (id)
    );`
	log.Println("<<<<<<<<..TABLE CREATED..>>>>>>>>>>")
	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

}

////////////////////////////////////////////////////////
func getTaskAll(db *sql.DB) []Task {
	rows, err := db.Query(`SELECT * FROM todo;`) // check err
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.Status)
		if err != nil {
			log.Fatal(err)
		} // check err
		tasks = append(tasks, t)
	}
	return tasks
}

////////////////////////////////////////////////////////
func insert(db *sql.DB, t Task) {
	rows, err := db.Query(`insert into todo(task,status) values(?,?)`, t.Name, t.Status) // check err
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

////////////////////////////////////////////////////////
func delete(db *sql.DB, t Task) {
	rows, err := db.Query(`DELETE FROM todo WHERE task=(?)`, t.Name) // check err
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

////////////////////////////////////////////////////////
func update(db *sql.DB, t Task) {
	rows, err := db.Query(`UPDATE todo set status=(?) WHERE task=(?)`, t.Status, t.Name) // check err
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}
