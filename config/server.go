// package config has all configurations of the server
package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents the Server configuration structure
type Server struct {
	Env                  string                    // application environment
	Port                 string                    // application port
	HostName             string                    // application host name
	DB                   *mongo.Client             // MongoDB cluster client
	Repos                *Repositories             // Repositories
}

type Repositories struct {
}
