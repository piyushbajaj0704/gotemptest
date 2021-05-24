package models

//TODO: should be moved to its own package
//TODO: models should also I be moved to its own package
import (
	"database/sql"
	"fmt"
)

// GetDB returns a db connection to our postgres db
// Can't defer the connection close here - ideas? Right now, will be deferred in the main
func GetDB() *sql.DB {
	fmt.Println("dbname:")

	// Get fast config
	configuration := GetConfig()

	psql := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", configuration.DbUsername, configuration.DbPassword, configuration.DbHost, configuration.DbPort, configuration.DbName)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}
	fmt.Println("created postgres conn:" + psql)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected")

	return db
	//testDb(db)
}

func testDb(db sql.DB) {
	fmt.Println("testing db..")
	dropSQL := "drop table if exists test;"
	db.Exec(dropSQL)
	createSQL := `create table test(id varchar(10));`
	db.Exec(createSQL)
	insertSQL := `insert into test(id) values ('12345jjj');`
	db.Exec(insertSQL)
	selectSQL := `select * from test;`
	row := db.QueryRow(selectSQL)

	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No test rows were returned!")
	case nil:
		fmt.Println(id)
	default:
		panic(err)
	}
	fmt.Println("done testing db..")
}
