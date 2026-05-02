/*
Copyright © 2026 Joel Faldín joelfaldin@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aero",
	Short: "A simple reverse proxy CLI tool made in go",
	Long: `
_aero_ is a CLI tool made in go.
Created to practice (and learn)
about proxies and reverse proxies.
(Currently under construction!)
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
