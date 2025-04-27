package player

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Player struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email      string             `json:"email" bson:"email"`
		Password   string             `json:"password" bson:"password"`
		Username   string             `json:"username" bson:"username"`
		CreateAt   time.Time          `json:"create_at" bson:"create_at"`
		UpdateAt   time.Time          `json:"update_at" bson:"update_at"`
		PlayerRole []PlayerRole       `bson:"player_role"`
	}

	PlayerRole struct {
		RoleTitle string `json:"role_tile" bson:"role_tile"`
		RoleCode  string `json:"role_code" bson:"role_code"`
	}

	PlayerProfileBson struct {
		Id       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email    string             `json:"email" bson:"email"`
		Username string             `json:"username" bson:"username"`
		CreateAt time.Time          `json:"create_at" bson:"create_at"`
		UpdateAt time.Time          `json:"update_at" bson:"update_at"`
	}
	PlayerSavingAccount struct {
		PlayerId string  `json:"player_id"`
		Balance  float64 `json:"balance" bson:"balance"`
	}
)
