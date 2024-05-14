package main

import ("net/http"
_ "github.com/mattn/go-sqlite3"
"encoding/json"
"log"
)

type Todo struct {
    Title string `json:"title"`
	Description string `json:"description"`
}


func main() {
	// create 'database'
	db := []Todo{} 

	person := Todo{Title:"Ballin",Description:"You have to ball"}

	db = append(db,person)

	// --http handler---
	http.HandleFunc("/todos/",func(w http.ResponseWriter, r *http.Request){
        ToDoListHandler(w, r,&db)
})
    // --http server at port8080--
	http.ListenAndServe(":8080", nil)
}

/*
======================== Request handler: maps requests to function ===============================
*/

func ToDoListHandler(w http.ResponseWriter, r *http.Request,db *[]Todo) {
	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method{
	case "GET":
		getTodos(w,db)
	case "POST":
		createTodo(w,r,db)
	default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
	// Your code here
}

/*
=============================== End of request handler =====================================
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
func getTodos(w http.ResponseWriter, db *[]Todo) {
	log.Println("====== Fetching data ======")
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(*db)
}

/*
========= function to create Todo  =========
*/
func createTodo(w http.ResponseWriter, r *http.Request,db *[]Todo) {
	log.Println("====== Posting data ======")
	var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// adds to db
    *db = append(*db,todo)
    w.WriteHeader(http.StatusCreated)
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


