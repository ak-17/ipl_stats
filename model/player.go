package model

type Player struct {
	Name      string  `json:"name"`
	Points    float64 `json:"points"`
	Catches   int     `json:"catches"`
	Matches   int     `json:"matches"`
	Team      string  `json:"team"`
	Stumpings int     `json:"stumpings"`
	Batting   Batting `json:"batting"`
	Bowling   Bowling `json:"bowling"`
}

type Batting struct {
	Innings    int     `json:"innings"`
	Runs       int     `json:"runs"`
	HighScore  int     `json:"high_score"`
	Fours      int     `json:"fours"`
	Sixes      int     `json:"sixes"`
	StrikeRate float64 `json:"strike_rate"`
	Hundreds   int     `json:"hundreds"`
	Fifties    int     `json:"fifties"`
	BallsFaced int     `json:"balls_faced"`
	Average    float64 `json:"avg"`
	NotOut     int     `json:"not_out"`
}

type Bowling struct {
	Innings        int     `json:"innings"`
	OversBowled    float64 `json:"overs_bowled"`
	FourWicketHaul int     `json:"4w"`
	FiveWicketHaul int     `json:"5w"`
	RunsConceded   int     `json:"runs_conceded"`
	StrikeRate     float64 `json:"strike_rate"`
	Wickets        int     `json:"wickets"`
	Average        float64 `json:"avg"`
	Economy        float64 `json:"economy"`
	Maidens        int     `json:"maidens"`
}
