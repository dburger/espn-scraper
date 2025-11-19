/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/dburger/espn-scraper/espn"
	"github.com/spf13/cobra"
)

// idsCmd represents the ids command
var idsCmd = &cobra.Command{
	Use:   "ids",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ids called")
		date := time.Date(2025, 10, 21, 0, 0, 0, 0, time.UTC)
		for i := 0; i < 200; i++ {
			gameIds, err := espn.FetchGameIds(date)
			if err != nil {
				panic(err)
			}
			for _, gameId := range gameIds {
				fmt.Println(date.Format("20060102"), gameId)
			}
			date = date.AddDate(0, 0, 1)
		}
	},
}

func init() {
	nbaCmd.AddCommand(idsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// idsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// idsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
