package main

import (
	"context"
	"errors"
	"github.com/Sigumaa/warp/db"
	"log"
	"net/url"
	"os"
)

var (
	ErrArgs = errors.New("two arguments are required. first is the path, second is the URI")
	ErrURI  = errors.New("please provide a valid URI")
)

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func parseArgs() (db.Link, error) {
	link := db.Link{}

	if len(os.Args) < 3 {
		return link, ErrArgs
	}
	if !isURL(os.Args[2]) {
		return link, ErrURI
	}

	link.Before = os.Args[1]
	link.After = os.Args[2]

	return link, nil
}

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
