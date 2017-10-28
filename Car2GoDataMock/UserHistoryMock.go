package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	// create database file
	fmt.Println("Start to create")

	dbAccessor := new(DatabaseAccessor)
	dbAccessor.CreateDatabase("historyDatabase")
	rand.Seed(time.Now().Unix())

	for i := 0; i < 1000000; i++ {
		row := new(HistoryDataRow)
		row.Time = GenerateRandomTime()
		row.LonCurrent, row.LatCurrent = GenerateRandomCoord()
		if rand.Intn(2) == 1 {
			row.Booked = true
			if rand.Intn(2) == 1 {
				row.LonDestination, row.LatDestination = GenerateRandomCoord()
			}
		} else {
			row.Booked = false
		}
		dbAccessor.AddRow(*row)
	}
	dbAccessor.Shutdown()
}

func GenerateRandomTime() time.Time {
	randomDay := rand.Intn(31 - 1) + 1
	randomHour := rand.Intn(24 - 1) + 1
	randomMin := rand.Intn(60 - 1) + 1
	randomSec := rand.Intn(60 - 1) + 1
	genTime := time.Date(2017, time.October,
		randomDay, randomHour,
		randomMin, randomSec,
		0, time.UTC)
	return genTime
}

func GenerateRandomCoord() (float64, float64) {
	randomLon := rand.Intn(40810000 - 40741549) + 40741549
	randomLat := rand.Intn((-73932326) - (-74006140)) + (-74006140)
	return float64(randomLon)/1000000, float64(randomLat)/1000000
}
