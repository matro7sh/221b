package cli

import (
	"github.com/spf13/cobra"

	"github.com/cmepw/221b/logger"
)

var debug bool

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "221b",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		logger.DebugMode = debug
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "activate debug mode")

	rootCmd.AddCommand(bake)
}
