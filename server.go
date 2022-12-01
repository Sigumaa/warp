package main

import (
	"context"
	"github.com/Sigumaa/warp/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func run(ctx context.Context, myDB *db.DB) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		log.Printf("GET / %s\n", c.Request().RemoteAddr)
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:path", func(c echo.Context) error {
		path := c.Param("path")
		link, err := myDB.GetLink(ctx, path)
		if err != nil {
			return c.String(http.StatusNotFound, "Not Found")
		}
		log.Printf("GET /%s %s\n", path, c.Request().RemoteAddr)
		log.Printf("redirect to %s\n", link.After)
		return c.Redirect(http.StatusMovedPermanently, link.After)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
