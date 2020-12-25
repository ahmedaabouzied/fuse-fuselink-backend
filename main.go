package main

import (
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/auth"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	userrepo "bitbucket.org/MoMoLab-dev/fuse.link-backend/user/repository"
	"context"
	"flag"
	"fmt"
	"os"
)

func main() {
	env := flag.String("env", "dev", "Running environment of the system. Can be either dev, stg, or prod")
	flag.Parse()
	ctx := context.Background()
	db, err := config.ConnectToMongoDB(ctx, os.Getenv("MONGODB_CONNECTION_STRING"))
	if err != nil {
		fmt.Println(os.Getenv("MONGODB_CONNECTION_STRING"))
		panic(err)
	}
	defer db.Disconnect(ctx)
	authHandler := &auth.JwtHandler{
		TokenKeyURI: os.Getenv("COGNITO_JWT_KEY_URL"),
	}
	userRepo := userrepo.NewUserRepository(db.Database(*env))
	repos := &config.Repositories{
		UserRepository: userRepo,
	}
	serverConfig := &config.Server{
		Repos:       repos,
		AuthHandler: authHandler,
	}
	startServer(serverConfig)
}
