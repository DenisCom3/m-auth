package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/model"
	"log"
)

func (s *serv) GetByName(ctx context.Context, name string) (*model.User, error) {

	var user *model.User

	userCache, err := s.cache.Get(ctx, fmt.Sprintf("user_%s", name))

	if err != nil {
		switch {
		case errors.Is(err, cache.ErrNotFound):
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

	user, err = s.userRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
