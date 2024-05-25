package user

import (
	"context"

	"github.com/rs/xid"
)

type Store interface {
	List(context.Context, func(*User)) error
	Get(context.Context, xid.ID, *User) error
	GetByEmail(context.Context, string, *User) error
	Insert(context.Context, *User) error
}
