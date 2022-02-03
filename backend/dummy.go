package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage!")
	fmt.Println("Endpoint hit!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/consumeapi", returnSearch) //pushing the consumer/vendor info to consumeapi
	// log.Fatal(http.ListenAndServe(":10000", nil))

	http.HandleFunc("/cityapi", returnCity) //pushing the cities json to cityapi
	log.Fatal(http.ListenAndServe(":10000", nil))
}

type search struct {
	ID       string `json:"ID"`
	Name     string `json: "Name"`
	Position string `json: "Position"`
}

type city struct {
	City_name string `json: "city_name"`
	City_id   int    `json: "city_id"`
}

var cities []city
var items []search

var id string
var name string
var pos string

var city_id int
var city_name string

//json for customer/vendor information

func returnSearch(w http.ResponseWriter, r *http.Request) { //can be revamped to push customer/vendor info

	fmt.Println("Returning the user search criteria:")
	json.NewEncoder(w).Encode(items)

}

//json for city information

func returnCity(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Returning the city search criteria:")
	json.NewEncoder(w).Encode(cities)

}

func main() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	createTable(sqliteDatabase)                                      // Create Database Tables

	createCity(sqliteDatabase)
	insertCity(sqliteDatabase, "Gainesville")

	// INSERT RECORDS
	insertStudent(sqliteDatabase, "0001", "Liana Kim", "Bachelor")
	insertStudent(sqliteDatabase, "0002", "Glen Rangel", "Bachelor")
	insertStudent(sqliteDatabase, "0003", "Martin Martins", "Master")
	insertStudent(sqliteDatabase, "0004", "Alayna Armitage", "PHD")
	insertStudent(sqliteDatabase, "0005", "Marni Benson", "Bachelor")
	insertStudent(sqliteDatabase, "0006", "Derrick Griffiths", "Master")
	insertStudent(sqliteDatabase, "0007", "Leigh Daly", "Bachelor")
	insertStudent(sqliteDatabase, "0008", "Marni Benson", "PHD")
	insertStudent(sqliteDatabase, "0009", "Klay Correa", "Bachelor")

	// DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)

	//printing city
	displayCity(sqliteDatabase)

	items = []search{
		{ID: id, Name: name, Position: pos},
	}

	cities = []city{
		{City_id: city_id, City_name: city_name},
	}

	handleRequests()
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createCity(db *sql.DB) {
	createCityTableSQL := `CREATE TABLE city (
		"city_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"city_name" TEXT	
	  );` // SQL Statement for Create Table

	log.Println("Create city table...")
	statement, err := db.Prepare(createCityTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("city table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertCity(db *sql.DB, city_name string) {
	log.Println("Inserting city record ...")
	insertCitySQL := `INSERT INTO city(city_name) VALUES (?)`
	statement, err := db.Prepare(insertCitySQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(city_name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {

	sqlStmt := `SELECT code, name, program FROM student WHERE code = $1;`

	row := db.QueryRow(sqlStmt, "0002")
	switch err := row.Scan(&id, &name, &pos); err {
	case sql.ErrNoRows:
		fmt.Println("No rows")
	case nil:
		fmt.Println(id, name, pos)
	default:
		panic(err)
	}
}

func displayCity(db *sql.DB) {

	sqlStmt := `SELECT city_ID, city_name FROM City WHERE city_ID = $1;`

	row := db.QueryRow(sqlStmt, "1")
	switch err := row.Scan(&city_id, &city_name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows")
	case nil:
		fmt.Println(city_id, city_name)
	default:
		panic(err)
	}

}
