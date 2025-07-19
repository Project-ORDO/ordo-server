package implementations

import (
	"context"
	"time"

	"github.com/Project-ORDO/ORDO-backEnd/config"
	models "github.com/Project-ORDO/ORDO-backEnd/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepo struct {
	collectionName string
}

func NewUserRepo() *userRepo {
	return &userRepo{
		collectionName: "users",
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = ptrTime(time.Now())
	user.IsActive = ptrBool(true)

	collection := config.GetCollection(r.collectionName)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *userRepo) UpdateUser(ctx context.Context, id primitive.ObjectID, updateData map[string]interface{}) error {
	collection := config.GetCollection(r.collectionName)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	collection := config.GetCollection(r.collectionName)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *userRepo) SoftDeleteUser(ctx context.Context, id primitive.ObjectID) error {
	collection := config.GetCollection(r.collectionName)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"isActive": false}})
	return err
}

func (r *userRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	collection := config.GetCollection(r.collectionName)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (r *userRepo) FindByName(ctx context.Context, name string) ([]models.User, error) {
	collection := config.GetCollection(r.collectionName)
	cursor, err := collection.Find(ctx, bson.M{"name": name})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err == nil {
			users = append(users, user)
		}
	}
	return users, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	collection := config.GetCollection(r.collectionName)
	var user models.User

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err // important: return nil here if not found
	}

	return &user, nil
}

// Helpers
func ptrBool(b bool) *bool          { return &b }
func ptrTime(t time.Time) *time.Time { return &t }
