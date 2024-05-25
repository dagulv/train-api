package main

import (
	"context"
	"log"
	"time"

	"github.com/dagulv/train-api/internal/adapter/http"
	"github.com/dagulv/train-api/internal/adapter/postgres"
	"github.com/dagulv/train-api/internal/domain/user"
	"github.com/dagulv/train-api/internal/env"
	"github.com/dagulv/train-api/internal/utils"
	"github.com/go-webauthn/webauthn/webauthn"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/xid"
)

func main() {
	ctx := context.Background()

	if err := start(ctx); err != nil {
		log.Fatal(err)
	}
}

func start(ctx context.Context) (err error) {
	env, err := env.GetEnv()

	if err != nil {
		return
	}

	db, err := postgres.Connect(ctx, env)

	if err != nil {
		return
	}

	defer db.Close()

	userStore := postgres.User(db)

	userService := user.Service{
		Store: userStore,
	}

	wconfig := &webauthn.Config{
		RPDisplayName: "Train",
		RPID:          "train.local",
		RPOrigins:     []string{"https://train.local"},
	}

	webAuthn, err := webauthn.New(wconfig)

	if err != nil {
		return
	}

	json := jsoniter.ConfigFastest

	server := http.Server{
		Json:     json,
		WebAuthn: webAuthn,
		User:     userService,
	}

	if err = createAdminUser(ctx, userService); err != nil {
		return
	}

	return server.StartServer(ctx)
}

func createAdminUser(ctx context.Context, userService user.Service) (err error) {
	var user user.User

	email := "admin@admin.admin"

	if err = userService.GetByEmail(ctx, email, &user); err == nil {
		return
	}

	user.Id = xid.NewWithTime(time.Now())
	user.FirstName = "Admin"
	user.LastName = "Admin"
	user.Email = email
	user.TimeCreated = utils.Timestamptz()
	user.TimeUpdated = user.TimeCreated

	return userService.Insert(ctx, &user)
}
