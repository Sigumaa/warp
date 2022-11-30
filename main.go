package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db := &DB{}
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer func(db *DB, ctx context.Context) {
		if err := db.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(db, ctx)

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
}
