package simulation

import (
	"fmt"
	"math/rand"
	"nba-scores/internal/models"
	"nba-scores/internal/store"
	"time"
)

const (
	Home = "home"
	Away = "away"
)

type SimulationService struct {
	db store.Store
}

func NewSimulatorService(db store.Store) SimulationService {
	return SimulationService{db: db}
}

// Simulation function simulates an NBA match
func (simulator *SimulationService) GetOneMinuteSimulation(minute int) []models.Match {
	matches := simulator.db.GetMatches()
	// we will simulate 1 min = 60 sec and max 24 sec as a attack time, attack way must be change every >24 second
	// we decided to start with home team
	for i, match := range matches {
		whoseBall := Home
		for i := 0; i < 60; i++ {
			i += generateRandomNumberFor24Second()
			isThree := generateRandomValueForThreePoint()
			match.AttackCount += 1
			match = updateStats(match, isThree, isSuccessful(), whoseBall)
			if whoseBall == Home {
				whoseBall = Away
			} else {
				whoseBall = Home
			}
		}
		match.MatchTime = fmt.Sprintf("%d:00", minute)
		matches[i] = match
		simulator.db.UpdateMatches(&matches)
	}
	return matches
}

func updateStats(match models.Match, three bool, successful bool, ball string) models.Match {
	if ball == Home {
		match.HomeScore += point(successful, three)
		if three {
			match.HomeStats.ThreePointAttempt += 1
			if successful {
				match.HomeStats.CompletedThreePointCount += 1
			}
			match.HomeStats.SuccessRateThreePoint = float32(match.HomeStats.CompletedThreePointCount) / float32(match.HomeStats.ThreePointAttempt)
		} else {
			match.HomeStats.TwoPointAttempt += 1
			if successful {
				match.HomeStats.CompletedTwoPointCount += 1
			}
			match.HomeStats.SuccessRateTwoPoint = float32(match.HomeStats.CompletedTwoPointCount) / float32(match.HomeStats.TwoPointAttempt)
		}
		match.HomeStats.AttackCount += 1
	} else {
		match.AwayScore += point(successful, three)
		if three {
			match.AwayStats.ThreePointAttempt += 1
			if successful {
				match.AwayStats.CompletedThreePointCount += 1
			}
			match.AwayStats.SuccessRateThreePoint = float32(match.AwayStats.CompletedThreePointCount) / float32(match.AwayStats.ThreePointAttempt)
		} else {
			match.AwayStats.TwoPointAttempt += 1
			if successful {
				match.AwayStats.CompletedTwoPointCount += 1
			}
			match.AwayStats.SuccessRateTwoPoint = float32(match.AwayStats.CompletedTwoPointCount) / float32(match.AwayStats.TwoPointAttempt)
		}
		match.AwayStats.AttackCount += 1
	}
	return match
}

func point(successfull bool, three bool) int {
	if !successfull {
		return 0
	}
	if three {
		return 3
	}
	return 2
}

func generateRandomNumberFor24Second() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 5
}

func isSuccessful() bool {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Float64() * 2 //just try to increase percentage
	if randomValue < 0.5 {
		return false
	}
	return true
}

func generateRandomValueForThreePoint() bool {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Float64()
	if randomValue < 0.5 {
		return false
	}
	return true
}
