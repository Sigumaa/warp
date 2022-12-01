package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Sigumaa/warp/db"
	"os"
	"strings"
)

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
	return nil
}
