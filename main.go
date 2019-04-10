package main

import (
	"context"
	"log"
	"os"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/csvmaker"
	"github.com/minhajuddinkhan/flutter_pk_firebase_export/fbase"
	"github.com/minhajuddinkhan/flutter_pk_firebase_export/repo"
)

func main() {
	ctx := context.TODO()
	privateKeyPath := "wtq-key.json"
	if len(os.Args) > 1 {
		privateKeyPath = os.Args[1]
	}

	client, err := fbase.NewFireStoreClient(ctx, privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repo.NewUserRepo(client.Collection("users"))
	users, err := userRepo.GetAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("result.csv")
	defer f.Close()

	csvMaker := csvmaker.New(f)
	if err := csvMaker.User(users); err != nil {
		log.Fatal(err)
	}

}
