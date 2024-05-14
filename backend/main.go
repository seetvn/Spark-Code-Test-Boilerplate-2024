package main

import ("net/http"
"database/sql"
_ "github.com/mattn/go-sqlite3"
"encoding/json"
"log"
 "os"
)

type Todo struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
	Description string `json:"description"`
	Priority string `json:"priority"`
}


func main() {

	//--create database connection--
	db, error := createConnection("todo_list.db")
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()
	log.Println("fuiyohhh")
	// --Create table --
    createTable := `
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            task TEXT,
			description TEXT,
			priority TEXT
        );
    `
	_, err := db.Exec(createTable)
    if err != nil {
		log.Println("PUIYOHHH")
        log.Fatal(err)
    }

	// --http handler---
	http.HandleFunc("/todos",func(w http.ResponseWriter, r *http.Request){
        ToDoListHandler(w, r,db)
})
    // --http server at port8080--
	http.ListenAndServe(":8080", nil)
	

}


func ToDoListHandler(w http.ResponseWriter, r *http.Request,db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method{
	case "GET":
		getTodos(w,db)
	case "POST":
		getTodos(w,db)
	}
	// Your code here
}


/*
--function to get Todos--
*/
func getTodos(w http.ResponseWriter, db *sql.DB) {
	log.Println("chamoy")

	// make query
    rows, err := db.Query("SELECT id, task, priority FROM todos")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

	// get all queries as array
    var todos []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Task, &todo.Description, &todo.Priority); err != nil {
            log.Fatal(err)
        }
        todos = append(todos, todo)
    }

	// return query via json
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}


/*
--function to create connection to db--
*/
func createConnection(dbpath string) (*sql.DB, error) {
	log.Println("---dbpath is " + dbpath + "----")

    // Check if the database file already exists
    _, err := os.Stat(dbpath)
    if os.IsNotExist(err) {
        // Database file does not exist, create it
        if _, err := os.Create(dbpath); err != nil {
            log.Fatal(err)
        }
        log.Println("Database file created.")
	}else{
		log.Println("---CHAMOY DATABASE PATH EXISTS ALREADY----")
	}

	// actual connection
	db, err := sql.Open("sqlite3", dbpath)
    	if err != nil {
        	log.Fatal(err)
			return nil, err
    	}
    	// defer db.Close()
	log.Println("Successfully connected db")
	return db,err
}

