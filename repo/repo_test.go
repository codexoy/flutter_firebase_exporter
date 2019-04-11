package repo

import (
	"context"
	"testing"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/fbase"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

	client, _ := fbase.NewFireStoreClient(context.TODO(), "../wtq-key.json")
	userRepo := NewUserRepo(client.Collection("users"))
	users, err := userRepo.GetAllUsers(context.TODO())
	assert.Nil(t, err)
	assert.NotEqual(t, len(users), 0)
	assert.NotEmpty(t, users[0].Email)
}
