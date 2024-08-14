package user

import (
	"context"
	"fmt"
	"log"

	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {

	var user *model.User

	userCache, err := s.cache.Get(ctx, fmt.Sprintf("user_%d", id))

	if err != nil {
		switch {
		case err == cache.ErrNotFound:
			log.Println("user not found in cache")
		default:
			return nil, err
		}
	}

	if user, ok := userCache.(*model.User); ok {
		return user, nil
	} else {
		log.Println("user not found in cache")
	}

	user, err = s.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
