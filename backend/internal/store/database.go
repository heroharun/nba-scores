package store

import (
	models "nba-scores/internal/models"
)

var store Store

// Create a in-memory database for teamDB
type TeamDatabase struct {
	Teams []models.Team
}

// Create a in-memory database for matchDB
type MatchDatabase struct {
	Matches []models.Match
}

type Store struct {
	teamDB  TeamDatabase
	matchDB MatchDatabase
}

func NewStore() (store *Store) {
	teams := &TeamDatabase{}
	teams.initTeamDatabase()
	matchs := &MatchDatabase{}
	matchs.initMatchDatabase()
	store = &Store{
		teamDB:  *teams,
		matchDB: *matchs,
	}
	return store
}

func (db *TeamDatabase) initTeamDatabase() {
	db.Teams = []models.Team{
		{ID: 1, Name: "Brooklyn Nets", Logo: "atlanta_hawks.png"},
		{ID: 2, Name: "Milwaukee Bucks", Logo: "boston_celtics.png"},
		{ID: 3, Name: "Philadelphia 76ers", Logo: "zart.png"},
		{ID: 4, Name: "Cleveland Cavaliers", Logo: "charlotte_hornets.png"},
		{ID: 5, Name: "Toronto Raptors", Logo: "chicago_bulls.png"},
		{ID: 6, Name: "Washington Wizards", Logo: "cleveland_cavaliers.png"},
		{ID: 7, Name: "Boston Celtics", Logo: "dallas_mavericks.png"},
		{ID: 8, Name: "Los Angeles Clippers", Logo: "denver_nuggets.png"},
		{ID: 9, Name: "Phoenix Suns", Logo: "detroit_pistons.png"},
		{ID: 10, Name: "Sacramento Kings", Logo: "golden_state_warriors.png"},
		{ID: 11, Name: "San Antonio Spurs", Logo: "houston_rockets.png"},
		{ID: 12, Name: "Oklahoma City Thunder", Logo: "indiana_pacers.png"},
		{ID: 13, Name: "Memphis Grizzlies", Logo: "los_angeles_clippers.png"},
		{ID: 14, Name: "Charlotte Hornets", Logo: "los_angeles_lakers.png"},
		{ID: 15, Name: "Miami Heat", Logo: "memphis_grizzlies.png"},
		{ID: 16, Name: "New York Knicks", Logo: "miami_heat.png"},
		{ID: 17, Name: "Orlando Magic", Logo: "orlando.png"},
		{ID: 18, Name: "Atlanta Hawks", Logo: "minnesota_timberwolves.png"},
		{ID: 19, Name: "Chicago Bulls", Logo: "new_orleans_pelicans.png"},
		{ID: 20, Name: "Dallas Mavericks", Logo: "new_york_knicks.png"},
		{ID: 21, Name: "Denver Nuggets", Logo: "oklahoma_city_thunder.png"},
		{ID: 22, Name: "Golden State Warriors", Logo: "orlando_magic.png"},
		{ID: 23, Name: "Houston Rockets", Logo: "philadelphia_76ers.png"},
		{ID: 24, Name: "Indiana Pacers", Logo: "phoenix_suns.png"},
		{ID: 25, Name: "Los Angeles Lakers", Logo: "portland_trail_blazers.png"},
		{ID: 26, Name: "Minnesota Timberwolves", Logo: "sacramento_kings.png"},
		{ID: 27, Name: "New Orleans Pelicans", Logo: "san_antonio_spurs.png"},
		{ID: 28, Name: "New York Knicks", Logo: "toronto_raptors.png"},
		{ID: 29, Name: "Portland Trail Blazers", Logo: "utah_jazz.png"},
		{ID: 30, Name: "Utah Jazz", Logo: "washington_wizards.png"},
	}
}

// initialize the match database only first week
func (db *MatchDatabase) initMatchDatabase() {
	db.Matches = []models.Match{
		{ID: 1, HomeTeamID: 1, AwayTeamID: 2, HomeTeam: "Brooklyn Nets", AwayTeam: "Milwaukee Bucks"},
		{ID: 2, HomeTeamID: 3, AwayTeamID: 4, HomeTeam: "Philadelphia 76ers", AwayTeam: "Cleveland Cavaliers", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 3, HomeTeamID: 5, AwayTeamID: 6, HomeTeam: "Toronto Raptors", AwayTeam: "Washington Wizards", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 4, HomeTeamID: 7, AwayTeamID: 8, HomeTeam: "Boston Celtics", AwayTeam: "Los Angeles Clippers", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 5, HomeTeamID: 9, AwayTeamID: 10, HomeTeam: "Phoenix Suns", AwayTeam: "Sacramento Kings", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 6, HomeTeamID: 11, AwayTeamID: 12, HomeTeam: "San Antonio Spurs", AwayTeam: "Oklahoma City Thunder", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 7, HomeTeamID: 13, AwayTeamID: 14, HomeTeam: "Memphis Grizzlies", AwayTeam: "Charlotte Hornets", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 8, HomeTeamID: 15, AwayTeamID: 16, HomeTeam: "Miami Heat", AwayTeam: "New York Knicks", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 9, HomeTeamID: 17, AwayTeamID: 18, HomeTeam: "Orlando Magic", AwayTeam: "Atlanta Hawks", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 10, HomeTeamID: 19, AwayTeamID: 20, HomeTeam: "Chicago Bulls", AwayTeam: "Dallas Mavericks", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 11, HomeTeamID: 21, AwayTeamID: 22, HomeTeam: "Denver Nuggets", AwayTeam: "Golden State Warriors", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 12, HomeTeamID: 23, AwayTeamID: 24, HomeTeam: "Houston Rockets", AwayTeam: "Indiana Pacers", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 13, HomeTeamID: 25, AwayTeamID: 26, HomeTeam: "Los Angeles Lakers", AwayTeam: "Minnesota Timberwolves", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 14, HomeTeamID: 27, AwayTeamID: 28, HomeTeam: "New Orleans Pelicans", AwayTeam: "New York Knicks", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
		{ID: 15, HomeTeamID: 29, AwayTeamID: 30, HomeTeam: "Portland Trail Blazers", AwayTeam: "Utah Jazz", HomeScore: 0,
			AwayScore:   0,
			MatchTime:   "",
			AttackCount: 0,
			HomeStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
			AwayStats: models.Stats{
				AttackCount:              0,
				TwoPointAttempt:          0,
				CompletedTwoPointCount:   0,
				ThreePointAttempt:        0,
				CompletedThreePointCount: 0,
				SuccessRateTwoPoint:      0,
				SuccessRateThreePoint:    0,
			},
		},
	}
}

// update batch matchDB
func (s *Store) UpdateMatches(matches *[]models.Match) {
	s.matchDB.Matches = *matches
}

func (s *Store) GetMatches() []models.Match {
	return s.matchDB.Matches
}
