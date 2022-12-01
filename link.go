package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/Sigumaa/warp/db"
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

func addLink(ctx context.Context, myDB *db.DB) error {
	// この辺めちゃくちゃ気持ち悪い、そもそも追加処理自体を別のツールにするべきかなぁ
	fmt.Println("do you want to add a link? (y/n)")
	var answer string
	bufio.NewScanner(os.Stdin).Scan()
	answer = bufio.NewScanner(os.Stdin).Text()

	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		for {
			link, err := parseArgs()
			if err != nil {
				return err
			}
			if err := myDB.AddLink(ctx, link); err != nil {
				return err
			}
			fmt.Println("link added.\ndo you want to add another link? (y/n)")
			bufio.NewScanner(os.Stdin).Scan()
			answer = bufio.NewScanner(os.Stdin).Text()
			if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
				continue
			} else {
				break
			}
		}
	}
	return nil
}
