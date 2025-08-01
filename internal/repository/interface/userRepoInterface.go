package interfaces

import (
	"context"

	models "github.com/Project-ORDO/ORDO-backEnd/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, id primitive.ObjectID, updateData map[string]interface{}) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	SoftDeleteUser(ctx context.Context, id primitive.ObjectID) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindByName(ctx context.Context, name string) ([]models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	SaveVerification(ctx context.Context, verification *models.UserVerification) error
	GetVerificationByToken(ctx context.Context, token string) (*models.UserVerification, error)
	DeleteVerificationByToken(ctx context.Context, token string) error
}
