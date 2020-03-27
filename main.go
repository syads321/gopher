package main

import (
	controller "eaciit/gopher/controller"
	types "eaciit/gopher/types"
	"encoding/json"
	"github.com/joho/godotenv"
	"net/http"
)

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t types.Query
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	result := controller.ExecuteQuery(t.Query, request)

	json.NewEncoder(rw).Encode(result)
}

func main() {
	godotenv.Load()
	http.HandleFunc("/graphql", parseGhPost)
	http.Handle("/", http.FileServer(http.Dir("asset")))
	http.HandleFunc("/upload", controller.UploadFile)
	http.HandleFunc("/route", controller.CalculateRoute)
	server := controller.SocketIO()
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.ListenAndServe(":8080", nil)
}
