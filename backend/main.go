package main

import ("net/http"
"database/sql"
_ "github.com/mattn/go-sqlite3"
"log"
 "os"
)

func main() {
	// Your code here
	/*
	--create database connection--
	*/
	create_connection("todo_list.db")
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Your code here
}

/*
--function to create connection to db--
*/
func create_connection(dbpath string) (*sql.DB, error) {
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
    	defer db.Close()
	log.Println("Successfully connected db")
	return db,err
}

