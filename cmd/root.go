/*
Copyright © 2025 heisenburg
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "espn-scraper",
	Short: "Tool for scraping statistical data from the ESPN website",
	Long: `A tool for scraping statistical data from the ESPN website.
It typically works by harvesting game IDs and then scraping individual
game pages for the data of interest.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.espn-scraper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
