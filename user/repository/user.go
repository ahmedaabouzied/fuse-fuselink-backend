package repository

import (
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/user"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) user.UserRepository {
	return &UserRepository{
		userCollection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	res, err := r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while creating user record")
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	updateParams := bson.D{{
		"$set", bson.D{
			{"username", user.Username},
			{"social_accounts", user.SocialAccounts},
		},
	}}
	res, err := r.userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, updateParams)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while updating user")
	}
	if res.MatchedCount != 1 {
		return nil, errors.New("repository error while updating user fields")
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, userID string) error {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.Wrap(err, "error parsing user ID")
	}
	res, err := r.userCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return errors.Wrap(err, "repository error while deleting user")
	}
	if res.DeletedCount != 1 {
		return errors.Wrap(err, "repository error while deleting user")
	}
	return nil
}
func (r *UserRepository) GetByID(ctx context.Context, ID string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing user ID")
	}
	res := r.userCollection.FindOne(ctx, bson.M{"_id": objectID})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "error finding user record")
	}
	var user entities.User
	err = res.Decode(&user)
	if err != nil {
		return nil, errors.Wrap(res.Err(), "error decoding returned result from DB")
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	res := r.userCollection.FindOne(ctx, bson.M{"username": username})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "error finding user record")
	}
	var user entities.User
	err := res.Decode(&user)
	if err != nil {
		return nil, errors.Wrap(res.Err(), "error decoding returned result from DB")
	}
	return &user, nil
}

func (r *UserRepository) GetByCognitoID(ctx context.Context, cognitoID string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	res := r.userCollection.FindOne(ctx, bson.M{"cognito_user_id": cognitoID})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "error finding user record")
	}
	var user entities.User
	err := res.Decode(&user)
	if err != nil {
		return nil, errors.Wrap(res.Err(), "error decoding returned result from DB")
	}
	return &user, nil
}
