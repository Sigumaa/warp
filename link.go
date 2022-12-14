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
	ErrURI = errors.New("please provide a valid URI")
)

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func addLink(ctx context.Context, myDB *db.DB) error {
	s := bufio.NewScanner(os.Stdin)
	// この辺めちゃくちゃ気持ち悪い、そもそも追加処理自体を別のツールにするべきかなぁ
	for {
		link := db.Link{}

		fmt.Println("please enter a short name for the link")
		s.Scan()
		link.Before = s.Text()

		fmt.Println("please enter a URI for the link")
		s.Scan()
		uri := s.Text()
		if !isURL(uri) {
			fmt.Println(ErrURI)
			fmt.Println("please continue...")
			continue
		}
		link.After = uri

		if err := myDB.AddLink(ctx, link); err != nil {
			fmt.Println(err)
			fmt.Println("please continue...")
			continue
		}
		fmt.Println("link added.\ndo you want to add another link? (y/N)")

		s.Scan()
		var answer string
		answer = s.Text()

		if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
			continue
		} else {
			break
		}
	}
	return nil
}
