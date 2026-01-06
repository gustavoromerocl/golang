package main

import (
	"log"

	"github.com/gustavoromerocl/social/cmd/internal/db"
	"github.com/gustavoromerocl/social/cmd/internal/env"
	"github.com/gustavoromerocl/social/cmd/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:pass@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")
	store := store.NewStorage(db)

	app := &application{
		store:  *store,
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
