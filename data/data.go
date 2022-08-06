/*
Copyright © 2022  <EMAIL ADDRESS>

*/
package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var db *sql.DB

func OpenDatabase() error {
	envErr := godotenv.Load("config/.env")
	if envErr != nil {
		log.Fatal(envErr)
		log.Fatal("Error loading env")
	}
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	database := os.Getenv("DBNAME")
	cfg := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, database,
	)
	var err error
	db, err = sql.Open("pgx", cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	pingStatus := db.Ping()
	if pingStatus != nil {
		log.Fatal(pingStatus)
	}
	fmt.Println("Connected!")
	return pingStatus
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS studybuddy (
		"idNote" SERIAL PRIMARY KEY ,
		"word" VARCHAR(50) NOT NULL,
		"definition" VARCHAR(100) NULL,
		"category" VARCHAR(40) NOT NULL
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Studybuddy table created")
}

func CreateNewNote(word string, definition string, category string) {
	insertNoteQuery := `INSERT INTO studybuddy(word, definition, category) VALUES ($1, $2, $3)`
	statement, err := db.Prepare(insertNoteQuery)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Study note created successfully.")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM studybuddy ORDER BY word")

	if err != nil {
		log.Fatalln(err)
	}
	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "—", definition)
	}
}
