package repo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/minhajuddinkhan/flutter_pk_firebase_export/models"
)

// UserRepo UserRepo
type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userrepo struct {
	collectionRef *firestore.CollectionRef
}

// NewUserRepo creates a new user repository
func NewUserRepo(cRef *firestore.CollectionRef) UserRepo {
	return &userrepo{cRef}
}

func (u *userrepo) GetAllUsers(ctx context.Context) ([]models.User, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context signalled done to return users")
	default:
		snapshot, err := u.collectionRef.Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}

		users := make([]models.User, len(snapshot))
		for j, snapUser := range snapshot {
			var user models.User
			if err := snapUser.DataTo(&user); err != nil {
				return nil, err
			}
			users[j] = user
		}
		return users, nil

	}
}
