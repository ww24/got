/**
 * fetch-image command
 *
 */

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/ww24/got/media"
)

// FetchImages command
var FetchImages = cli.Command{
	Name:      "fetch-images",
	Usage:     "fetch user images",
	Aliases:   []string{"fi"},
	ArgsUsage: "<screen_name>",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "size",
			Value: 10,
			Usage: "size of fetch image",
		},
		cli.IntFlag{
			Name:  "lastid",
			Value: 0,
			Usage: "fetch start tweet ID",
		},
		cli.BoolFlag{
			Name:  "orig",
			Usage: "set original image flag",
		},
		cli.BoolFlag{
			Name:  "download",
			Usage: "auto download images",
		},
	},
	Action: func(ctx *cli.Context) {
		if len(ctx.Args()) < 1 {
			cli.ShowCommandHelp(ctx, "fetch-images")
			os.Exit(1)
		}
		screenName := ctx.Args()[0]

		urls := make([]string, 0, 100)

		urls, err := media.GetImageUrls(screenName, ctx.Int("size"), uint64(ctx.Int("lastid")))
		if err != nil {
			panic(err)
		}

		if ctx.Bool("orig") {
			for i := range urls {
				urls[i] += ":orig"
			}
		}

		if ctx.Bool("download") {
			for _, url := range urls {
				fmt.Println(url)
				media.Download(url)
			}
		} else {
			fmt.Println(strings.Join(urls, "\n"))
		}

		log.Printf("Get %d urls.\n", len(urls))
	},
}
