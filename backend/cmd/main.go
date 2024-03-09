package main

import (
	httpserver "nba-scores/cmd/http"
	"nba-scores/internal/api"
	"nba-scores/internal/simulation"
	"nba-scores/internal/store"
)

func main() {
	//initialize the store, teams and matches
	db := store.NewStore()

	//initialize simulator service
	simulatorSvc := simulation.NewSimulatorService(*db)

	//initialize the API
	api := api.New(simulatorSvc)

	startHttpServer(api)

}

func startHttpServer(api *api.API) {
	router := httpserver.Router(*api)
	router.Start()
}
