package authHandler

import (
	"context"

	"github.com/bunkai888/hello-microservice/module/auth"
	authPb "github.com/bunkai888/hello-microservice/module/auth/authPb"
	"github.com/bunkai888/hello-microservice/module/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.CredentialSearchReq) (*auth.CredentialRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
