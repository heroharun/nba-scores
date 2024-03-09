package models

// Match struct represents a simulated NBA match
// Match struct represents a simulated NBA match
type Match struct {
	ID          int    `json:"id"`
	HomeTeamID  int    `json:"home_team_id"`
	AwayTeamID  int    `json:"away_team_id"`
	HomeTeam    string `json:"home_team"`
	AwayTeam    string `json:"away_team"`
	HomeScore   int    `json:"home_score"`
	AwayScore   int    `json:"away_score"`
	MatchTime   string `json:"match_time"`
	AttackCount int    `json:"attack_count"`
	HomeStats   Stats  `json:"home_stats"`
	AwayStats   Stats  `json:"away_stats"`
}

type Stats struct {
	AttackCount              int     `json:"attack_count"`
	TwoPointAttempt          int     `json:"two_point_attempt"`
	CompletedTwoPointCount   int     `json:"completed_two_point_count"`
	ThreePointAttempt        int     `json:"three_point_attempt"`
	CompletedThreePointCount int     `json:"completed_three_point_count"`
	SuccessRateTwoPoint      float32 `json:"success_rate_two_point"`
	SuccessRateThreePoint    float32 `json:"success_rate_three_point"`
}
