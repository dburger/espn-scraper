/*
Copyright © 2025 heisenburg
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
	Short: "Dumps largest lead stats for the given game IDs",
	Long: `Dumps the largest lead stats scraped from the ESPN website
for the given game IDs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		idfile, err := cmd.Flags().GetString("idfile")
		if err != nil {
			return err
		}
		for gameId := range espn.Fileids(idfile) {
			stat, err := espn.FetchLargestLead(gameId)
			if err == nil {
				fmt.Println(gameId, stat)
			} else {
				fmt.Println(gameId, err)
			}
		}

		return nil
	},
}

func init() {
	nbaCmd.AddCommand(leadsCmd)
	leadsCmd.Flags().StringP("idfile", "i", "", "File containing game ids")
	err := leadsCmd.MarkFlagRequired("idfile")
	if err != nil {
		panic(err)
	}
}
