package cmd

import (
	"fmt"
	"os"

	"github.com/Aloe-Corporation/muuidwizard/cmd/decode"
	"github.com/Aloe-Corporation/muuidwizard/cmd/encode"
	"github.com/Aloe-Corporation/muuidwizard/cmd/generate"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "muuidwizard",
	Short: "Decode and encode uuid from mongo compass UUID representation",
}

func Execute() {
	encode.Init(&rootCmd)
	decode.Init(&rootCmd)
	generate.Init(&rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
