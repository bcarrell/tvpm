package main

import (
	tdb "github.com/bcarrell/tvpm/db"
	"github.com/bcarrell/tvpm/episode"
	"github.com/bcarrell/tvpm/indexer"
	"github.com/bcarrell/tvpm/series"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	// create tables
	tdb.CreateTables()
	about := About()
	tvpm := cli.NewApp()
	tvpm.Name = about["name"]
	tvpm.Usage = about["usage"]
	tvpm.Version = about["version"]
	tvpm.Flags = []cli.Flag{
		cli.StringFlag{
			"quality",
			"720",
			"default video quality [sd, hd, 720, 1080]",
		},
	}
	tvpm.Commands = []cli.Command{
		{
			Name:      "find",
			ShortName: "f",
			Usage:     "find an episode",
			Action:    episode.FindEpisode,
		},
		{
			Name:      "find-series",
			ShortName: "fs",
			Usage:     "find a series",
			Action:    series.FindSeries,
		},
		// {
		// 	Name:      "list",
		// 	ShortName: "l",
		// 	Usage:     "list installed episodes by series",
		// 	Action:    command.List,
		// },
		{
			Name:      "add-indexer",
			ShortName: "ai",
			Usage:     "add a newznab indexer",
			Flags: []cli.Flag{
				cli.StringFlag{
					"apiKey",
					"",
					"your API key for the indexer",
				},
			},
			Action: indexer.AddIndexer,
		},
		// {
		// 	Name:      "add-series",
		// 	ShortName: "as",
		// 	Usage:     "add a series",
		// 	Action:    command.AddSeries,
		// },
	}
	tvpm.Run(os.Args)
}
