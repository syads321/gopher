package controller

import (
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

type UpdatedRoute struct {
	TripID string
	Lat    float64
	Lng    float64
}

func SocketIO() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnEvent("/update-route", "join", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		ReceivedData := &UpdatedRoute{}
		json.Unmarshal([]byte(msg), ReceivedData)
		server.JoinRoom(ReceivedData.TripID, s)
		fmt.Println("joined room:" + s.ID())
		return "joined room"
	})

	server.OnEvent("/update-route", "route", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		ReceivedData := &UpdatedRoute{}
		json.Unmarshal([]byte(msg), ReceivedData)
		UpdatedRoute, _ := json.Marshal(ReceivedData)
		UpdatedRouteStr := string(UpdatedRoute)
		server.BroadcastToRoom(ReceivedData.TripID, "route", UpdatedRouteStr)
		return "received"
	})
	return server
}
