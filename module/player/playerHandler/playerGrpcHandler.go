package playerHandler

import (
	"context"

	playerPb "github.com/bunkai888/hello-microservice/module/player/playerPb"
	"github.com/bunkai888/hello-microservice/module/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{playerUsecase: playerUsecase}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
