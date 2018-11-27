package database

import (
	"database/sql"
	"fmt"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/environment"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

// Connection ...
var Connection *sql.DB

// Connect ...
func Connect() {
	var datasource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", environment.DATABASEUSERNAME, environment.DATABASEPASSWORD, environment.DATABASEHOST, environment.DATABASEPORT, environment.DATABASENAME)
	conn, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	Connection = conn
}
