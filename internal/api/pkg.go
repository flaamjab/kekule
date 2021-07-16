package api

import (
	"errors"
	"os"

	"github.com/flaamjab/kekule/internal/db"
)

func Run(addr string) {
	r := router()

	if _, err := os.Stat(db.DB_PATH); errors.Is(err, os.ErrNotExist) {
		db.Initialize()
	}

	r.Run(addr)
}
