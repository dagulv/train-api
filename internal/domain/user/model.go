package user

import (
	"github.com/jackc/pgx/v5/pgtype"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/xid"
)

type User struct {
	Id          xid.ID             `json:"id"`
	FirstName   string             `json:"firstName" validate:"required"`
	LastName    string             `json:"lastName"`
	Email       string             `json:"email" validate:"required"`
	PublicKey   string             `json:"publicKey"`
	TimeCreated pgtype.Timestamptz `json:"timeCreated"`
	TimeUpdated pgtype.Timestamptz `json:"timeUpdated"`
}

func (u *User) EncodeToStream(s *jsoniter.Stream) {
	s.WriteObjectField("id")
	s.WriteString(u.Id.String())

	s.WriteMore()
	s.WriteObjectField("firstName")
	s.WriteString(u.FirstName)

	s.WriteMore()
	s.WriteObjectField("lastName")
	s.WriteString(u.LastName)

	s.WriteMore()
	s.WriteObjectField("email")
	s.WriteString(u.Email)

	s.WriteMore()
	s.WriteObjectField("publicKey")
	s.WriteString(u.PublicKey)

	s.WriteMore()
	s.WriteObjectField("timeCreated")
	s.WriteString(u.TimeCreated.Time.String())

	s.WriteMore()
	s.WriteObjectField("timeUpdated")
	s.WriteString(u.TimeUpdated.Time.String())
}
