package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// 短縮URL用。
// pathに指定された任意の文字列がDBに存在する場合、リダイレクトする。
// なければ、404を返す。

type Link struct {
	Before string
	After  string
}

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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:path", func(c echo.Context) error {
		path := c.Param("path")
		link, err := db.GetLink(ctx, path)
		if err != nil {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.Redirect(http.StatusMovedPermanently, link.After)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
