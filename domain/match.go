package domain

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchArea struct{
	ID 					int 				`json:"id"`
	Name				string 				`json:"name"`
	CountryCode			string 				`json:"code"`
	FlagURL				string 				`json:"flag"`
}

type Matchcompetition struct{
	ID 					int 				`json:"id"`
	Name				string 				`json:"name"`
	Code				string 				`json:"code"`
	Type 				string 				`json:"type"`
	Emleem				string 				`json:"emblem"`
}

type MatchSeason struct {
	ID 					int 				`json:"id"`
	StartDate 			string 				`json:"startDate"`
	EndDate 			string 				`json:"endDate"`
	CurrentMatchDay 	int 				`json:"currentMatchDay"`
	Winner 				*string				`json:"winner"`
}

type Teams struct{
	ID 					int 				`json:"id"`
	Name 				string 				`json:"name"`
	ShortName 			string 				`json:"shortName"`
	TLA 				string 				`json:"tla"`
	Crest 				string 				`json:"crest"`
}

type TimeScore struct {
	Home 				int 				`json:"home"`
	Away 				int 				`json:"away"`
}

type Score struct {
	Winner 				*string				`json:"winner"`
	Duration 			string 				`json:"duration"`
	FullTime 			TimeScore			`json:"fullTime"`
	HalfTime 			TimeScore			`json:"halfTime"`
}

type Refree struct {
	ID 					int 				`json:"id"`
	Name 				string 				`json:"name"`
	Type 				string 				`json:"type"`
	Nationality 		string 				`json:"nationality"`
}

type Match struct{
	Area				MatchArea 			`json:"area"`
	Competition			Matchcompetition 	`json:"competition"`
	Season				MatchSeason 		`json:"season"`
	ID 					string 				`json:"id"`
	UTCDate				string 				`json:"utcDate"`
	Status 				string  			`json:"status"`
	MatchDay			int 				`json:"matchday"`
	Stage 				string  			`json:"stage"`
	Group 				*string				`json:"group"`
	LastUpdated			string 				`json:"lastUpdated"`
	HomeTeam 			Teams 				`json:"homeTeam"`
	AwayTeam 			Teams 				`json:"awayTeam"`
	Score 				Score				`json:"score"`
	Refree 				Refree				`json:"refree"`
}

type MatchResponse struct{
	Area				MatchArea 			`json:"area"`
	Competition			Matchcompetition 	`json:"competition"`
	Season				MatchSeason 		`json:"season"`
	ID 					string 				`json:"id"`
	UTCDate				string 				`json:"utcDate"`
	Status 				string  			`json:"status"`
	MatchDay			int 				`json:"matchday"`
	Stage 				string  			`json:"stage"`
	Group 				*string				`json:"group"`
	LastUpdated			string 				`json:"lastUpdated"`
	HomeTeam 			Teams 				`json:"homeTeam"`
	AwayTeam 			Teams 				`json:"awayTeam"`
	Score 				Score				`json:"score"`
	Refree 				Refree				`json:"refree"`
}

type MatchesUseCase interface{
	GetTeamMatches(c *gin.Context, collection *mongo.Collection, limit int, offset int,teamID string) ([]MatchResponse, error)
}