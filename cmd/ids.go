/*
Copyright © 2025 heisenburg
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
	Short: "Dumps the game IDs for the current NBA season",
	Long: `Dumps the game IDs for the current NBA season. Currently
this is hardcoded to the 2025 season. Note that this tool does not,
by default, know how to weed out the NBA cup final and various all-star
weekend events.`,
	Run: func(cmd *cobra.Command, args []string) {
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
}
