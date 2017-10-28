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

type EnvData struct {
    monthday int
    weekday int
		Hour int
		IsPriceHigher bool
		DoesPolicyApply bool
		IsRush bool
		IsGoodWeather bool
    IsFallout bool
    IsWeekend bool
    EventID int
}

type DatabaseAccessor struct {
    databaseHistory *sql.DB
    //databaseEnv *sql.DB
}

func (accessor *DatabaseAccessor) OpenDB(databasePath string) {
    database, err := sql.Open("sqlite3", databasePath)
    if err != nil {
        fmt.Println("Error occured while creating database")
        return
    }
    accessor.databaseHistory = database
}

func (accessor *DatabaseAccessor) CreateDatabaseHistory() bool {
    statement, _ :=  accessor.databaseHistory.Prepare(`CREATE TABLE IF NOT EXISTS HistoryData
         (id INTEGER PRIMARY KEY, time INTEGER, lonCurrent REAL,
         latCurrent REAL, lonDestination REAL , latDestination REAL, booked BOOL)`)
    statement.Exec()

    return true
}

func (accessor *DatabaseAccessor) AddRowToHistory(data HistoryDataRow) bool {
    statement, err := accessor.databaseHistory.Prepare(`INSERT INTO HistoryData
        (time, lonCurrent, latCurrent, lonDestination, latDestination, booked) VALUES (?, ?, ?, ?, ?, ?)`)
    if err != nil {
        fmt.Println("Error AddRowToHistory prep" + err.Error())
        return false
    }
    _, err = statement.Exec(int(data.Time.Unix()), data.LonCurrent, data.LatCurrent,
        data.LonDestination, data.LatDestination, data.Booked)
    if err != nil {
        fmt.Println("Error AddRowToHistory exec" + err.Error())
        return false
    }
    return true
}

func (accessor *DatabaseAccessor) CreateDatabaseEnv() bool {
    statement, err :=  accessor.databaseHistory.Prepare(`CREATE TABLE IF NOT EXISTS EnvData
         (id INTEGER PRIMARY KEY, monthDay INTEGER, weekday INTEGER, hour INTEGER, isRush BOOL,
         isGoodWeather BOOL, isFallout BOOL, isWeekend BOOL, IsPriceHigher BOOL, doesPolicyApply BOOL, EventID INTEGER)`)
         if err != nil {
            fmt.Println("Error CreateDatabaseEnv pr" + err.Error())
            return false
        }
    statement.Exec()

    return true
}

func (accessor *DatabaseAccessor) AddRowToEnv(data EnvData) bool {
    statement, err :=  accessor.databaseHistory.Prepare(`INSERT INTO EnvData
        (monthDay, weekday, hour, isRush, isGoodWeather,
					isFallout, isWeekend, IsPriceHigher, doesPolicyApply, EventID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
    if err != nil {
        fmt.Println("Error AddRowToEnv prep" + err.Error())
        return false
    }
    _, err = statement.Exec(data.monthday, data.weekday, data.Hour, data.IsRush,
        data.IsGoodWeather, data.IsFallout, data.IsWeekend, data.IsPriceHigher, data.DoesPolicyApply, data.EventID)
    if err != nil {
        fmt.Println("Error AddRowToEnv exec" + err.Error())
        return false
    }
    return true
}

func (accessor *DatabaseAccessor) Shutdown() bool {
    accessor.databaseHistory.Close()
    return true
}
