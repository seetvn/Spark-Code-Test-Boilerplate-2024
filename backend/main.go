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
	Finished bool `json:"finished"`
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
			priority TEXT,
			finished INTEGER
        );
    `
	_, err := db.Exec(createTable)
    if err != nil {
		log.Println("PUIYOHHH")
        log.Fatal(err)
    }

	// --http handler---
	http.HandleFunc("/todos/",func(w http.ResponseWriter, r *http.Request){
        ToDoListHandler(w, r,db)
})
    // --http server at port8080--
	http.ListenAndServe(":8080", nil)
	

}

/*
=========== Request handler: maps requests to function ===================
*/
func ToDoListHandler(w http.ResponseWriter, r *http.Request,db *sql.DB) {
	//CORS
	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method{
	case "GET":
		getTodos(w,db)
	case "POST":
		createTodo(w,r,db)
	case "PATCH":
		updateTodo(w,r,db)
	case "PUT":
		updateTodo(w,r,db)
	case "DELETE":
		deleteTodo(w,r,db)
	
	default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	// Your code here
}
/*
============= End of request handler ==================
*/


/* 
=======================================================
*/


/*
--------------------------------------------------------------------
---------------------------------------------------------------------
------------------------START OF CRUD OPERATIONS------------------------
------------------------------------------------------------------------
---------------------------------------------------------------------
*/

/*
======== function to get Todos ===========
*/
func getTodos(w http.ResponseWriter, db *sql.DB) {
	log.Println("chamoy")

	// make query
    rows, err := db.Query("SELECT id, task, description, priority, finished FROM todos")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

	// get all queries as array
    var todos []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Task, &todo.Description, &todo.Priority, &todo.Finished); err != nil {
            log.Fatal(err)
        }
        todos = append(todos, todo)
    }

	// return query via json
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

/*
========= function to create Todo  =========
*/
func createTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("INSERT INTO todos (task, description, priority,finished) VALUES (?, ?, ?, ?)", todo.Task, todo.Description, todo.Priority, todo.Finished)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusCreated)
}

/*
========= function to PATCH/PUT Todo =============
*/
func updateTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("===hello====")
    // Extract todo ID from the request URL
    todoID := r.URL.Path[len("/todos/"):]
	log.Println(todoID)
    
    // Parse JSON request body into Todo struct
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Execute SQL UPDATE statement to update todo in the database
    _, err := db.Exec("UPDATE todos SET task = ?, description = ?, priority = ? WHERE id = ?", todo.Task, todo.Description, todo.Priority, todoID)
    if err != nil {
        log.Fatal(err)
    }
    
    w.WriteHeader(http.StatusOK)
}

/*
========= function to DELETE Todo =============
*/
func deleteTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    // Extract todo ID from the request URL
    todoID := r.URL.Path[len("/todos/"):]
	log.Println("====="+todoID+"========")
    
    // Execute SQL DELETE statement to delete todo from the database
    _, err := db.Exec("DELETE FROM todos WHERE id = ?", todoID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}


/*
--------------------------------------------------------------------
---------------------------------------------------------------------
------------------------END OF CRUD OPERATIONS------------------------
------------------------------------------------------------------------
---------------------------------------------------------------------
*/


/*
=========================================================================
*/



/*
--------------------------------------------------------------------
---------------------------------------------------------------------
------------------------CREATE CONNECTION ------------------------
------------------------------------------------------------------------
---------------------------------------------------------------------
*/
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
/*
--------------------------------------------------------------------
---------------------------------------------------------------------
------------------------END OF CREATE CONNECTION ------------------------
------------------------------------------------------------------------
---------------------------------------------------------------------
*/

