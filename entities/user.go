package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username"`
	SocialAccounts []SocialAccount    `bson:"social_accounts" `
}
