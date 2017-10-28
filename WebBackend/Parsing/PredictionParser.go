package Parsing

import (
	 "os"
	 "log"
	 "encoding/csv"
	 "strconv"
   )

type ZoneDemand struct {
	Hour int
	ZoneID int
	Demands int
}

type PredictionParser struct {
	ParsedData []ZoneDemand
}

func (parser *PredictionParser) ParseInputFile() {
	file, err := os.Open("predict.csv")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	r := csv.NewReader(file)
	
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to read file")
	}

	for _, record := range records {
		hour, _ := strconv.Atoi(record[0])
		zoneID, _ := strconv.Atoi(record[1])
		demands, _ := strconv.Atoi(record[2])
		tmp := ZoneDemand{hour, zoneID, demands}
		parser.ParsedData = append(parser.ParsedData, tmp)
	}
}