package playerHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewHttpPlayerHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase}
}
