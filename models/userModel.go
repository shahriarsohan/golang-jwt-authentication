package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required , min=2,max=100"`
	LastName     *string            `json:"last_name"`
	Password     *string            `json:"password" validate:"required , min=4,max=14"`
	Email        *string            `json:"email"`
	Phone        *string            `json:"phone"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" validate:"required eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token"`
	CreateAt     time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}
