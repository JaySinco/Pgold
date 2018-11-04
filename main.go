package main

import (
	"log"
	"os"

	"github.com/jaysinco/Pgold/control"
	"github.com/jaysinco/Pgold/market"
	"github.com/jaysinco/Pgold/pg"
	"github.com/jaysinco/Pgold/policy"
	"github.com/jaysinco/Pgold/server"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "pgold"
	app.Usage = "ICBC paper gold trader assist system"
	app.Flags = []cli.Flag{
		pg.ConfigFlag,
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "market",
			Usage:  "Crawl market data into database continuously",
			Action: pg.Setup(market.Run),
		},
		cli.Command{
			Name:  "export",
			Usage: "Export market data from database into file",
			Flags: []cli.Flag{
				pg.OutfileFlag,
				pg.StartDateFlag,
				pg.EndDateFlag,
				pg.OnlyTxOpenFlag,
			},
			Action: pg.Setup(market.Export),
		},
		cli.Command{
			Name:  "import",
			Usage: "Import market data from file into database",
			Flags: []cli.Flag{
				pg.InfileFlag,
				pg.StartDateFlag,
				pg.EndDateFlag,
				pg.OnlyTxOpenFlag,
			},
			Action: pg.Setup(market.Import),
		},
		cli.Command{
			Name:   "autosave",
			Usage:  "Autosave market data into file daily",
			Action: pg.Setup(market.Autosave),
		},
		cli.Command{
			Name:   "server",
			Usage:  "Run http server showing market history data ",
			Action: pg.Setup(server.Run),
		},
		cli.Command{
			Name:   "realtime",
			Usage:  "Email trade tips continuously based on policy",
			Action: pg.Setup(policy.Realtime),
		},
		cli.Command{
			Name:  "test",
			Usage: "Loopback test strategy using history data",
			Flags: []cli.Flag{
				pg.PolicyFlag,
				pg.StartDateFlag,
				pg.EndDateFlag,
			},
			Action: pg.Setup(policy.Test),
		},
		cli.Command{
			Name:  "multitask",
			Usage: "Run serveral tasks simultaneously",
			Flags: []cli.Flag{
				pg.TaskSetFlag,
			},
			Action: pg.Setup(control.MutltiTask),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("[ERROR] pgold: %v", err)
	}
}
