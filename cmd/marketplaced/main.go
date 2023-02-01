package main

import (
	"os"

	"github.com/OmniFlix/marketplace/app"
	"github.com/OmniFlix/marketplace/cmd/marketplaced/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "MARKETPLACED", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
