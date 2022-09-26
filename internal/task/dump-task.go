package task

import (
	"database/sql"
	"fmt"
	"kirky/devy/internal/constants"
	"log"
	"os"

	"github.com/JamesStewy/go-mysqldump"
	_ "github.com/go-sql-driver/mysql"
)

// ========== Extract Schema ==========
func DumpDB(schema string) {
	// Open connection to database
	log.Print("Connecting to database...")

	// Used constants file to avoid dependency on .env file
	username := constants.DB_USERNAME
	password := constants.DB_PASSWORD
	hostname := constants.DB_HOST
	port := constants.DB_PORT
	dbname := schema

	dumpDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Established connection.")

	dumpFilenameFormat := fmt.Sprintf("%s-2006-01-02T150405", dbname)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return
	}

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		log.Fatal("Error registering databse:", err)
		return
	}

	// Dump database to file
	resultFilename, err := dumper.Dump()
	if err != nil {
		log.Fatal("Error dumping:", err)
		return
	}
	log.Printf("File is saved to %s", resultFilename)

	dumper.Close()
}
