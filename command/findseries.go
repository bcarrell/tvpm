package command

import (
	"fmt"
	"github.com/bcarrell/tvpm/api"
	"github.com/codegangsta/cli"
	"os"
	"text/tabwriter"
)

func FindSeries(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Wrong number of arguments.\n\n")
		cli.ShowCommandHelp(c, "find-series")
		os.Exit(1)
	}
	seriesQ := args[0]
	data := api.FindSeries(seriesQ)
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)
	for _, item := range data {
		fmt.Fprintf(
			w,
			"%s\t%s\t%d\t%s\n",
			item.Title,
			item.Network,
			item.Year,
			item.Url,
		)
	}
	w.Flush()
}
