// package config has all configurations of the server
package config

import (
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/auth"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents the Server configuration structure
type Server struct {
	Env         string           // application environment
	Port        string           // application port
	HostName    string           // application host name
	DB          *mongo.Client    // MongoDB cluster client
	Repos       *Repositories    // Repositories
	AuthHandler auth.AuthHandler // Authentication Handler
	UserUsecase user.UserUsecase
}

type Repositories struct {
	UserRepository user.UserRepository
}
