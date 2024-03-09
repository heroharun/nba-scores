package api

import "nba-scores/internal/simulation"

type API struct {
	Simulator simulation.SimulationService
}

func New(simulator simulation.SimulationService) *API {
	return &API{
		Simulator: simulator,
	}
}
