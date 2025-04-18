package playerHandler

import "github.com/bunkai888/hello-microservice/module/player/playerUsecase"

type (
	playerGrpcHandlerService struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandlerService {
	return &playerGrpcHandlerService{playerUsecase}
}
