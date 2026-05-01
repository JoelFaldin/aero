/*
Copyright © 2026 Joel Faldín joelfaldin@gmail.com
*/
package cmd

import (
	"aero/internal/app"
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the basic command of aero",
	Long: `Start the basic functionality of the tool.
	Remember to start the different go servers first!`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println("error processing flag:", err)
		}

		if flag {
			app.Handler(true)
		} else {
			app.Handler(false)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Local flags:
	runCmd.Flags().Bool("verbose", false, "Add more output of the proxy for debugging.")
}
