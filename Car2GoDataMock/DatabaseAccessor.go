package main

import (
	"database/sql"
    "fmt"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

type HistoryDataRow struct {
    Time time.Time
    LonCurrent float64
    LatCurrent float64
    LonDestination float64
    LatDestination float64
    Booked bool
}

type DatabaseAccessor struct {
	database *sql.DB
}

func (accessor *DatabaseAccessor) CreateDatabase(databasePath string) bool {
    database, err := sql.Open("sqlite3", databasePath)
    if err != nil {
        fmt.Println("Error occured while creating database")
        return false
    }
    accessor.database = database
    statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS HistoryData
         (id INTEGER PRIMARY KEY, time INTEGER, lonCurrent REAL,
         latCurrent REAL, lonDestination REAL , latDestination REAL, booked BOOL)`)
    statement.Exec()

    return true
}

func (accessor *DatabaseAccessor) AddRow(data HistoryDataRow) bool {
    statement, err := accessor.database.Prepare(`INSERT INTO HistoryData 
        (time, lonCurrent, latCurrent, lonDestination, latDestination, booked) VALUES (?, ?, ?, ?, ?, ?)`)
    if err != nil {
        fmt.Println("Error AddRow prep" + err.Error())
        return false
    }
    _, err = statement.Exec(int(data.Time.Unix()), data.LonCurrent, data.LatCurrent,
        data.LonDestination, data.LatDestination, data.Booked)
    if err != nil {
        fmt.Println("Error AddRow exec" + err.Error())
        return false
    }
    return true
}

func (accessor *DatabaseAccessor) Shutdown() bool {
    accessor.database.Close() 
    return true
}