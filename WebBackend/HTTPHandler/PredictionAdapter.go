package HTTPHandler

import (
	"math/rand"
	"net/http"
	"encoding/json"
	"time"
	"log"
)

type DemandResponse struct {
	DemandsByZone []int `json:"_demandsByZone"`
	DemandsByRatio []float64 `json:"_demandsByRatio"`
	ZoneIDs []int `json:"_zoneIDs"`
}

type PredictionAdapter struct {
	
}

type UserDemandData struct {
	Demands int
	ZoneID int
}

func (adapter *PredictionAdapter) HandleDemandsRequest(w http.ResponseWriter, day int, hour int) bool {
	
	demandResponse := &DemandResponse{}
/*	var demandsByZone []int
	var demandsByRatio []float64
	var zoneIDs []int*/
	rand.Seed(time.Now().Unix())
	for i:=0; i <=7; i++ {
		randomDemands := rand.Intn(100 - 50) + 50
		demandResponse.DemandsByZone = append(demandResponse.DemandsByZone, randomDemands)
		demandResponse.ZoneIDs = append(demandResponse.ZoneIDs, i)
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