package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UUID         string             `bson:"uuid" json:"uuid"`
	Email        string             `bson:"email" json:"email" validate:"required,email"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Password     string             `bson:"password" json:"password" validate:"required"`
	Phone        *string            `bson:"phone,omitempty" json:"phone,omitempty"`
	ProfileImage *string            `bson:"profileImage,omitempty" json:"profileImage,omitempty"`
	AuthProvider *[]string          `bson:"authProvider,omitempty" json:"authProvider,omitempty"`
	IsActive     *bool              `bson:"isActive,omitempty" json:"isActive,omitempty"`
	CreatedAt    *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	LastLogin    *time.Time         `bson:"lastLogin,omitempty" json:"lastLogin,omitempty"`
	Batch        *string            `bson:"batch,omitempty" json:"batch,omitempty"`
}
