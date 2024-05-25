package postgres

import (
	"context"

	"github.com/dagulv/train-api/internal/domain/user"
	"github.com/rs/xid"
	"github.com/webmafia/pg"
)

type userStore struct {
	db *pg.DB
}

func User(pool *pg.DB) user.Store {
	return userStore{
		db: pool,
	}
}

func (s userStore) List(ctx context.Context, cb func(user *user.User)) (err error) {
	rows, err := s.db.Query(ctx, `
		SELECT
			"id",
			"firstName",
			"lastName",
			"email",
			"publicKey",
			"timeCreated",
			"timeUpdated"
		FROM %T
	`, Users)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user user.User

		if err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PublicKey, &user.TimeCreated, &user.TimeUpdated); err != nil {
			return
		}

		cb(&user)
	}

	return
}

func (s userStore) Get(ctx context.Context, userId xid.ID, user *user.User) (err error) {
	row := s.db.QueryRow(ctx, `
		SELECT
			"id",
			"firstName",
			"lastName",
			"email",
			"publicKey",
			"timeCreated",
			"timeUpdated"
		FROM %T
		WHERE id = %s
	`, Users, userId)

	return row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PublicKey, &user.TimeCreated, &user.TimeUpdated)
}

func (s userStore) GetByEmail(ctx context.Context, email string, user *user.User) (err error) {
	row := s.db.QueryRow(ctx, `
		SELECT
			"id",
			"firstName",
			"lastName",
			"email",
			"publicKey",
			"timeCreated",
			"timeUpdated"
		FROM %T
		WHERE email = %s
	`, Users, email)

	return row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PublicKey, &user.TimeCreated, &user.TimeUpdated)
}

func (s userStore) Insert(ctx context.Context, user *user.User) (err error) {
	vals := s.db.AcquireValues()
	defer s.db.ReleaseValues(vals)

	vals.
		Value("id", user.Id).
		Value("firstName", user.FirstName).
		Value("lastName", user.LastName).
		Value("email", user.Email).
		Value("publicKey", user.PublicKey).
		Value("timeCreated", user.TimeCreated).
		Value("timeUpdated", user.TimeUpdated)

	_, err = s.db.InsertValues(ctx, Users, vals)

	return
}
