package util

import "github.com/ak-17/ipl_stats/model"

func GetFantasyPoints(player model.Player) float64 {
	points := float64(player.Matches) * 4
	points = points + 8*float64(player.Catches) + 12*float64(player.Stumpings)
	points = points + getBattingFantasyPoints(player)
	points = points + getBowlingFantasyPoints(player)
	return points
}

func getBowlingFantasyPoints(player model.Player) float64 {
	points := float64(player.Bowling.Wickets)*25 + 8*float64(player.Bowling.FourWicketHaul) + 16*float64(player.Bowling.FiveWicketHaul)
	points = points + float64(player.Bowling.Maidens)*8
	return points
}

func getBattingFantasyPoints(player model.Player) float64 {
	points := float64(player.Batting.Runs) + float64(player.Batting.Fours) + 2*float64(player.Batting.Sixes)
	points = points + 8*float64(player.Batting.Fifties) + 16*float64(player.Batting.Hundreds)
	return points
}
