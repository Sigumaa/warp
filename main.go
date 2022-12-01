package main

import (
	"context"
	"github.com/Sigumaa/warp/db"
	"log"
)

func main() {
	ctx := context.TODO()

	myDB := &db.DB{}
	if err := myDB.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer func(myDB *db.DB, ctx context.Context) {
		if err := myDB.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(myDB, ctx)

	if err := myDB.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	if err := addLink(ctx, myDB); err != nil {
		log.Fatal(err)
	}

	run(ctx, myDB)
}
