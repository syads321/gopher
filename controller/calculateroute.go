package controller

import (
	"context"
	ctx "eaciit/gopher/dbcontext"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type RouteResponse struct {
	FormatVersion string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func CalculateRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	path := "https://api.tomtom.com/routing/1/calculateRoute/52.3679,4.8786:52.3679,4.8798/json?instructionsType=coded&computeBestOrder=true&computeTravelTimeFor=all&report=effectiveSettings&routeType=fastest&avoid=unpavedRoads&travelMode=truck&vehicleCommercial=true&key=Q8ZbhC6idTbdjVvKa2IAgUJvqplianeo"
	resp, err := myClient.Get(path)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	routeResp := make(map[string]interface{})

	json.NewDecoder(resp.Body).Decode(&routeResp)
	//	mapstructure.Decode(routeResp, &routeResp)
	routes := routeResp["routes"].([]interface{})
	// for _, leg := range legs {
	// 	fmt.Println(leg)
	// }
	//	summary := routes["summary"].(interface{})
	var content bson.D
	var layout = "2006-01-02T15:04:05+01:00"
	for _, m := range routes {
		route := m.(map[string]interface{})
		summary := route["summary"].(map[string]interface{})
		departureTime := summary["departureTime"].(string)

		depart, err := time.Parse(layout, departureTime)
		if err != nil {
			fmt.Println("error parsing depart time")
		}
		arrivalTime := summary["arrivalTime"].(string)
		arrival, err := time.Parse(layout, arrivalTime)
		if err != nil {
			fmt.Println("error parsing depart time")
		}
		//	LiveTrafficIncidentsTravelTimeInSeconds
		content = bson.D{
			bson.E{Key: "LengthInMeters", Value: summary["lengthInMeters"].(float64)},
			bson.E{Key: "TravelTimeInSeconds", Value: summary["travelTimeInSeconds"].(float64)},
			bson.E{Key: "TrafficDelayInSeconds", Value: summary["trafficDelayInSeconds"].(float64)},
			bson.E{Key: "DepartureTime", Value: depart},
			bson.E{Key: "ArrivalTime", Value: arrival},
			bson.E{Key: "NoTrafficTravelTimeInSeconds", Value: summary["noTrafficTravelTimeInSeconds"].(float64)},
			bson.E{Key: "HistoricTrafficTravelTimeInSeconds", Value: summary["historicTrafficTravelTimeInSeconds"].(float64)},
			bson.E{Key: "LiveTrafficIncidentsTravelTimeInSeconds", Value: summary["liveTrafficIncidentsTravelTimeInSeconds"].(float64)},
		}
	}

	_, err = ctx.MainContext("Route").InsertOne(context.TODO(), content)
	if err != nil {
		fmt.Println("error when insert")
	}
	return
}
