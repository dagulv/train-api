package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Timestamptz(ta ...time.Time) pgtype.Timestamptz {
	var t time.Time

	if len(ta) == 0 || ta[0].IsZero() {
		t = time.Now()
	} else {
		t = ta[0]
	}

	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
