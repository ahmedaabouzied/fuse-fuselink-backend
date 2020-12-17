package main

import (
	"context"
	"flag"
	"fmt"
    "bitbucket.org/MoMoLab-dev/fuse-backend/config"
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
}
