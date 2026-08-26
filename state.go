package main

import (
	"github.com/kaszta1274/gator/internal/config"
	"github.com/kaszta1274/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
