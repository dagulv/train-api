package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dagulv/train-api/internal/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}

}

func run(_ context.Context, args []string) (err error) {
	env, err := env.GetEnv()

	if err != nil {
		return
	}

	dir, err := migrationsPath()

	if err != nil {
		return
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return errors.New("migration files not found")
	}

	m, err := migrate.New("file://"+dir, env.DatabaseUrl)

	if err != nil {
		return
	}

	defer m.Close()

	if len(args) > 0 {
		switch args[0] {
		case "up":
			var steps int

			if len(args) > 1 {
				steps, _ = strconv.Atoi(args[1])
			}

			if steps <= 0 {
				steps = 1
			}

			return m.Steps(steps)

		case "down":
			var steps int

			if len(args) > 1 {
				steps, _ = strconv.Atoi(args[1])
			}

			if steps <= 0 {
				steps = 1
			}

			return m.Steps(-steps)

		case "force":
			var v int

			if len(args) > 1 {
				v, _ = strconv.Atoi(args[1])
			}

			if v <= 0 {
				return
			}

			return m.Force(v)
		}
	}

	return m.Up()
}

func migrationsPath() (dir string, err error) {

	// If go.mod is found
	if dir, err = findDir("go.mod"); err == nil {
		dir = filepath.Join(dir, "internal", "adapter", "postgres", "migrations")
		return
	}

	// Otherwise, migrations might exist alongside the executable
	dir, err = os.Executable()

	if err != nil {
		return
	}

	dir = filepath.Join(filepath.Dir(dir), "migrations")

	return
}

func findDir(filename string) (dir string, err error) {
	dir, err = os.Getwd()

	if err != nil {
		return
	}

	for dir != "." {
		if _, err = os.Stat(filepath.Join(dir, filename)); os.IsNotExist(err) {
			dir = filepath.Dir(dir)
			continue
		}

		return
	}

	err = fmt.Errorf("no %s file found", filename)

	return
}
