package playerUsecase

import "github.com/bunkai888/hello-microservice/module/player/playerRepository"

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository: playerRepository}
}
