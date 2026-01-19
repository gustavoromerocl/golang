package main

import (
	"github.com/gustavoromerocl/social/cmd/internal/db"
	"github.com/gustavoromerocl/social/cmd/internal/env"
	"github.com/gustavoromerocl/social/cmd/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://user:pass@localhost:5432/social?sslmode=disable")
	// This is just a placeholder main function.
	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
