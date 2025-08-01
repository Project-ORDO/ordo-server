package models

import "time"

type UserVerification struct {
	Email        string     `bson:"email"`
	Name         string     `bson:"name"`
	Password     string     `bson:"password"`
	Token        string     `bson:"token"`
	AttemptCount int        `bson:"attemptCount"`
	ExpiresAt    time.Time  `bson:"expiresAt"`
	CreatedAt    time.Time  `bson:"createdAt"`
	VerifiedAt   *time.Time `bson:"verifiedAt,omitempty"`
}
