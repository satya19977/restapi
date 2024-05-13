package db
 
import (
    "database/sql"
    "fmt"
    "io/ioutil"
    _ "modernc.org/sqlite"
)
 
var DB *sql.DB
 
func InitDB() {
 
    db, err := sql.Open("sqlite", "api.sql")
    if err != nil {
        panic("Database could not connect: " + err.Error())
    }
 
    DB = db

    // Run migration
    err = runMigration()
    if err != nil {
        panic("Could not run migration: " + err.Error())
    }

    fmt.Println("Migration executed successfully!")
 
}
 
func runMigration() error {
    // Read migration SQL file
    migrationSQL, err := ioutil.ReadFile("migration.sql")
    if err != nil {
        return err
    }

    // Execute migration script
    _, err = DB.Exec(string(migrationSQL))
    if err != nil {
        return err
    }
    return nil

}
