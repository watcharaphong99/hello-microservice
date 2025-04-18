package server

import (
	"github.com/bunkai888/hello-microservice/module/player/playerHandler"
	"github.com/bunkai888/hello-microservice/module/player/playerRepository"
	"github.com/bunkai888/hello-microservice/module/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewHttpPlayerHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	player.GET("", s.healthCheckService)
}
