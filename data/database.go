package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

func New() {
	// Read the config file
	data, err := ioutil.ReadFile(config.GetConfig().GetString("Apps.database.config"))
	if err != nil {
		panic(err)
	}

	// Decode the YAML data into a Config struct
	var config config.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// Create a database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Apps.Database.User, config.Apps.Database.Pass,
		config.Apps.Database.IP, config.Apps.Database.Port,
		config.Apps.Database.Name))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Perform a test query
	rows, err := db.Query("SELECT * FROM table_name")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the query results
	for rows.Next() {
		var column1 string
		var column2 int
		err = rows.Scan(&column1, &column2)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Column 1: %s, Column 2: %d\n", column1, column2)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// Output a success message
	fmt.Println("Database connection test successful!")

}
