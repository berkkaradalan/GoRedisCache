package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamArea struct {
	ID 								int						`bson:"id"`
	Name							string					`bson:"name"`
	CountryCode						string					`bson:"code"`
	FlagUrl							*string					`bson:"flag"`

}
type TeamSquad struct {
	ID 								int						`bson:"id"`
	Name							string					`bson:"name"`
	Position						string					`bson:"position"`
	DateOfBirth						string					`bson:"dateOfBirth"`
	Nationality						string					`bson:"nationality"`
}

type TeamCompetition struct {
	ID								int						`bson:"id"`
	Name							string					`bson:"name"`
	Code							string					`bson:"code"`
	Type							string					`bson:"type"`
	EmblemUrl						*string					`bson:"emblemUrl"`
}

type TeamCoachContract struct {
	From							string					`bson:"from"`
	Until							string					`bson:"until"`
}

type TeamCoach struct {
	ID								int						`bson:"id"`
	FirstName						string					`bson:"firstname"`
	LastName						string					`bson:"lastname"`
	Name							string					`bson:"name"`
	DateOfBirth						string					`bson:"dateOfBirth"`
	Nationality						string					`bson:"nationality"`
	Contract 						*TeamCoachContract		`bson:"contract"`
}

type Team struct {
	Area 							*TeamArea				`bson:"area"`
	ID								int						`bson:"id"`
	Name							string					`bson:"name"`
	ShortName						string					`bson:"shortName"`
	Tla								string					`bson:"tla"`
	CrestUrl						*string					`bson:"crestUrl"`
	Address							string					`bson:"address"`
	Website							string					`bson:"website"`
	Founded							int						`bson:"founded"`
	ClubColors						string					`bson:"clubColors"`
	Venue							string					`bson:"venue"`
	Competition						[]*TeamCompetition		`bson:"competition"`
	Squad							[]*TeamSquad			`bson:"squad"`
	LastUpdated						*string					`bson:"lastUpdated"`
}

type TeamResponse struct {
	Area 							*TeamArea				`bson:"area"`
	ID								int						`bson:"id"`
	Name							string					`bson:"name"`
	ShortName						string					`bson:"shortName"`
	Tla								string					`bson:"tla"`
	CrestUrl						*string					`bson:"crestUrl"`
	Address							string					`bson:"address"`
	Website							string					`bson:"website"`
	Founded							int						`bson:"founded"`
	ClubColors						string					`bson:"clubColors"`
	Venue							string					`bson:"venue"`
	Competition						[]*TeamCompetition		`bson:"competition"`
	Squad							[]*TeamSquad			`bson:"squad"`
	LastUpdated						*string					`bson:"lastUpdated"`
}

type TeamsUseCase interface {
	GetTeamById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client,team_id string) (TeamResponse, error)
	GetTeams (ctx *gin.Context, collection *mongo.Collection,limit int, offset int) ([]TeamResponse, error)
}