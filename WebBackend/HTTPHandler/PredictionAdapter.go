package HTTPHandler

import (
	"net/http"
	"encoding/json"
	"log"
	"DratutiTeam/WebBackend/Parsing"
)

type DemandResponse struct {
	DemandsByZone []int `json:"_demandsByZone"`
	DemandsByRatio []float64 `json:"_demandsByRatio"`
	ZoneIDs []int `json:"_zoneIDs"`
}

type PredictionAdapter struct {
	Parser *Parsing.PredictionParser
}

type UserDemandData struct {
	Demands int
	ZoneID int
}

func (adapter *PredictionAdapter) InitParser() {
	adapter.Parser = new(Parsing.PredictionParser)
	adapter.Parser.ParseInputFile()
}

func (adapter *PredictionAdapter) HandleDemandsRequest(w http.ResponseWriter, hour int) bool {
	
	demandResponse := &DemandResponse{}
/*	var demandsByZone []int
	var demandsByRatio []float64
	var zoneIDs []int*/

	for _, value := range adapter.Parser.ParsedData {
		if hour == value.Hour {
			demandResponse.DemandsByZone = append(demandResponse.DemandsByZone, value.Demands)
			demandResponse.ZoneIDs = append(demandResponse.ZoneIDs, value.ZoneID)
		}
	}

	maxValue:=0
	for _, value := range demandResponse.DemandsByZone {
		if value >= maxValue {
			maxValue = value
		}
	}

	for _, value := range demandResponse.DemandsByZone {
		demandResponse.DemandsByRatio = append(demandResponse.DemandsByRatio, float64(value)/float64(maxValue))
	}

	response, errM := json.Marshal(&demandResponse)
	if errM != nil {
		log.Println(errM)
		return false
	}
	w.Write(response)

	return true
}