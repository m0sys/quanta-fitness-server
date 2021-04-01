package userstore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

const (
	projectId string = "quanta-fitness"
	collName  string = "Users"
)

type fireStore struct{}

func NewFirestoreUserStore() UserStore {
	return &fireStore{}
}

func (*fireStore) Save(user entity.BaseUser) (entity.User, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return entity.User{}, formatErr("failed to connect to firestore", err)
	}

	// newUser := entity.User{
	// 	BaseUser: user,
	// }

	defer client.Close()
	_, _, err2 := client.Collection(collName).Add(ctx, map[string]interface{}{
		"username":   user.Username,
		"email":      user.Email,
		"hashed_pwd": user.Password,
		"joined":     time.Now(),
		"updated_at": time.Now(),
		"weight":     0.0,
		"height":     0.0,
		"gender":     "N/A",
	})
	if err2 != nil {
		return entity.User{}, formatErr("failed to write to firestore", err2)
	}

	return entity.User{}, nil
}

func (*fireStore) FindUserByUsername(username string) (entity.User, bool, error) {
	return entity.User{}, false, nil
}

func (*fireStore) FindUserByEmail(email string) (entity.User, bool, error) {
	return entity.User{}, false, nil
}

func formatErr(msg string, err error) error {
	return fmt.Errorf("%s: %s: %w", "userstore", msg, err)
}
