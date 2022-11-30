package main

import (
	"context"
	"github.com/Sigumaa/warp/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:path", func(c echo.Context) error {
		path := c.Param("path")
		link, err := myDB.GetLink(ctx, path)
		if err != nil {
			return c.String(http.StatusNotFound, "Not found")
		}
		return c.Redirect(http.StatusMovedPermanently, link.After)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
