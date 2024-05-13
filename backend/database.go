import ("fmt")

func create_connection(dbpath string) {
	fmt.Println("chamoy " + dbpath)
    // db, err := sql.Open("sqlite3", "your_database.db")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer db.Close()

    // // Check if the database file already exists
    // _, err = os.Stat("your_database.db")
    // if os.IsNotExist(err) {
    //     // Database file does not exist, create it
    //     if _, err := os.Create("your_database.db"); err != nil {
    //         log.Fatal(err)
    //     }
    //     log.Println("Database file created.")
    // }

    // Use the db connection here...
}
