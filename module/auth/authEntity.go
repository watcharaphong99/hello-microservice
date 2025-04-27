package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		PlayerId    string             `json:"player_id" bson:"player_id"`
		RoleCode    int                `json:"role_code" bson:"role_code"`
		AccessToken string             `json:"access_token" bson:"acess_token"`
		CreateAt    time.Time          `json:"create_at" bson:"create_at"`
		UpdateAt    time.Time          `json:"update_at" bson:"update_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title string             `json:"title" bson:"title"`
		Code  int                `json:"code" bson:"code"`
	}
)
