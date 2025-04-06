package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerContract struct{
	Start						string								`bson:"start"`
	Until						string								`bson:"until"`
}

type PlayerTeamArea struct{
	ID 							int									`bson:"id"`
	Name						string								`bson:"name"`
	CountryCode					string								`bson:"code"`
	FlagUrl						*string								`bson:"flagUrl"`
}

type PlayerTeamRunninCompetitions struct{
	ID							int									`bson:"id"`
	Name						string								`bson:"name"`
	Code						string								`bson:"code"`
	Type						string								`bson:"type"`
	EmblemUrl					*string								`bson:"emblemUrl"`
}

type PlayerTeam struct {
	Area 							*PlayerTeamArea					`bson:"area"`
	ID 								int								`bson:"id"`
	Name							string							`bson:"name"`
	ShortName						string							`bson:"shortName"`
	Tla								string							`bson:"tla"`
	CrestUrl						*string							`bson:"crestUrl"`
	Address							*string							`bson:"address"`
	Website							*string							`bson:"website"`
	Founded							int								`bson:"founded"`
	ClubColors						string							`bson:"clubColors"`
	Venue							string							`bson:"venue"`
	Contract						*PlayerContract					`bson:"contract"`
	Competitions					[]*PlayerTeamRunninCompetitions	`bson:"competitions"`
}

type Player struct {
	ID 								int								`bson:"id"`
	Name							string							`bson:"name"`
	FirstName						string							`bson:"firstName"`
	LastName						string							`bson:"lastName"`
	DateOfBirth						string							`bson:"dateOfBirth"`
	Nationality						string							`bson:"nationality"`
	Section							string							`bson:"section"`
	Position						string							`bson:"position"`
	ShirtNumber						int								`bson:"shirtNumber"`
	LastUpdated						*string							`bson:"lastUpdated"`
}

type PlayerResponse struct {
	ID 								int								`bson:"id"`
	Name							string							`bson:"name"`
	FirstName						string							`bson:"firstName"`
	LastName						string							`bson:"lastName"`
	DateOfBirth						string							`bson:"dateOfBirth"`
	Nationality						string							`bson:"nationality"`
	Section							string							`bson:"section"`
	Position						string							`bson:"position"`
	ShirtNumber						int								`bson:"shirtNumber"`
	LastUpdated						*string							`bson:"lastUpdated"`
	CurrentTeam						*PlayerTeam						`bson:"team"`
}

type PlayerUseCase interface {
	GetPlayerById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client,player_id string) (PlayerResponse, error)
}