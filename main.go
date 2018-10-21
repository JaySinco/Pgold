package main

import (
	"log"
	"os"

	"github.com/jaysinco/Pgold/control"
	"github.com/jaysinco/Pgold/hint"
	"github.com/jaysinco/Pgold/market"
	"github.com/jaysinco/Pgold/show"
	"github.com/jaysinco/Pgold/utils"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "pgold"
	app.Usage = "ICBC paper gold trader helping system"
	app.Flags = []cli.Flag{
		utils.ConfigFlag,
	}
	app.Commands = []cli.Command{
		market.MarketCmd,
		market.ExportCmd,
		market.ImportCmd,
		show.ShowCmd,
		hint.HintCmd,
		control.BatchCmd,
	}

	app.After = func(c *cli.Context) error {
		if utils.DB != nil {
			return utils.DB.Close()
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("run app: %v", err)
	}
}
