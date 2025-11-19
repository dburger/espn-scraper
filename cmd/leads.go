/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dburger/espn-scraper/espn"
	"github.com/spf13/cobra"
)

// leadsCmd represents the leads command
var leadsCmd = &cobra.Command{
	Use:   "leads",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("leads called")
		for gameId := range espn.Fileids("2025ids.txt") {
			stat, err := espn.FetchLargestLead(gameId)
			if err == nil {
				fmt.Println(gameId, stat)
			} else {
				fmt.Println(gameId, err)
			}
		}
	},
}

func init() {
	nbaCmd.AddCommand(leadsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leadsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leadsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
