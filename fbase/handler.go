package fbase

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewFireStoreClient creates a new firestore from secret key
func NewFireStoreClient(ctx context.Context, filepath string) (*firestore.Client, error) {

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("firestore creation context signalled done")
	default:
		if _, err := os.Stat(filepath); err != nil {
			return nil, err
		}
		clientOptions := option.WithCredentialsFile(filepath)

		app, err := firebase.NewApp(ctx, nil, clientOptions)
		if err != nil {
			return nil, fmt.Errorf("cannot create new firebase app. err: %v", err)
		}
		return app.Firestore(ctx)
	}

}
