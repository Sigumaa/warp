package main

import (
	"context"
	"fmt"
	"github.com/Sigumaa/warp/db"
	"log"
	"net/url"
	"os"
	"time"
)

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	myDB := &db.DB{}

	if err := myDB.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer func(myDB *db.DB, ctx context.Context) {
		if err := myDB.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(myDB, ctx)

	link := db.Link{}

	if len(os.Args) < 3 {
		fmt.Println("two arguments are required. first is the path, second is the url.")
		return
	}
	if !isURL(os.Args[2]) {
		fmt.Println("Please provide a valid URL.")
		return
	}

	link.Before = os.Args[1]
	link.After = os.Args[2]

	if err := myDB.AddLink(ctx, link); err != nil {
		log.Fatal(err)
	}
}
