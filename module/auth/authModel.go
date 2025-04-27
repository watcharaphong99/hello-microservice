package auth

import (
	"time"

	"github.com/bunkai888/hello-microservice/module/player"
)

type (
	PlayerLoginReq struct {
		Email    string `json:"email" form:"email" validate:"require,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenReq struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"require.max=500"`
	}

	InserPlayerRole struct {
		PlayerId string `json:"player_id" validate:"required"`
		RoleCode int    `json:"role_id"`
	}

	ProfileIntercepter struct {
		*player.PlayerProfile
		Credential *Credential `json:"credential"`
	}

	CredentialRes struct {
		Id          string    `json:"_id" bson:"_id,omitempty"`
		PlayerId    string    `json:"player_id" bson:"player_id"`
		RoleCode    int       `json:"role_code" bson:"role_code"`
		AccessToken string    `json:"access_token" bson:"acess_token"`
		CreateAt    time.Time `json:"create_at"`
		UpdateAt    time.Time `json:"update_at"`
	}
)
