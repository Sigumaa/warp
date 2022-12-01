package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Sigumaa/warp/db"
	"log"
	"os"
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

	s := bufio.NewScanner(os.Stdin)
	fmt.Println("run server or add link? (s/a)")
	s.Scan()
	switch s.Text() {
	case "s":
		run(ctx, myDB)
	case "a":
		if err := addLink(ctx, myDB); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid answer")
	}
}
