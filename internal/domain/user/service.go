package user

import (
	"context"
	"time"

	"github.com/dagulv/train-api/internal/utils"
	"github.com/rs/xid"
)

type Service struct {
	Store Store
}

func (s Service) List(ctx context.Context, cb func(user *User)) (err error) {
	return s.Store.List(ctx, cb)
}

func (s Service) Get(ctx context.Context, userId xid.ID, user *User) (err error) {
	return s.Store.Get(ctx, userId, user)
}

func (s Service) GetByEmail(ctx context.Context, email string, user *User) (err error) {
	return s.Store.GetByEmail(ctx, email, user)
}

func (s Service) Insert(ctx context.Context, user *User) (err error) {
	user.Id = xid.NewWithTime(time.Now())
	user.TimeCreated = utils.Timestamptz()
	user.TimeUpdated = user.TimeCreated

	return s.Store.Insert(ctx, user)
}
