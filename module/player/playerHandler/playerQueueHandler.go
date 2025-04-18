package playerHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase}
}
