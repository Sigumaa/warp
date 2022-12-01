package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/Sigumaa/warp/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	// この辺めちゃくちゃ気持ち悪い、そもそも追加処理自体を別のツールにするべきかなぁ
	fmt.Println("do you want to add a link? (y/n)")
	var answer string
	bufio.NewScanner(os.Stdin).Scan()
	answer = bufio.NewScanner(os.Stdin).Text()
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		for {
			link, err := parseArgs()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err := myDB.AddLink(ctx, link); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("link added")

			fmt.Println("do you want to add another link? (y/n)")
			bufio.NewScanner(os.Stdin).Scan()
			answer = bufio.NewScanner(os.Stdin).Text()
			if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
				continue
			} else {
				break
			}
		}
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:path", func(c echo.Context) error {
		path := c.Param("path")
		link, err := myDB.GetLink(ctx, path)
		if err != nil {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.Redirect(http.StatusMovedPermanently, link.After)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
