/*
Copyright © 2025 heisenburg
*/
package cmd

import (
	"fmt"

	"github.com/dburger/espn-scraper/espn"
	"github.com/spf13/cobra"
)

// comebacksCmd represents the comebacks command
var comebacksCmd = &cobra.Command{
	Use:   "comebacks",
	Short: "Dumps comeback stats for the given game IDs",
	Long: `Dumps the comeback stats scraped from the ESPN website
for the given game IDs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		idfile, err := cmd.Flags().GetString("idfile")
		if err != nil {
			return err
		}

		for gameId := range espn.Fileids(idfile) {
			stat, err := espn.FetchComeback(gameId)
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
	nbaCmd.AddCommand(comebacksCmd)
	comebacksCmd.Flags().StringP("idfile", "i", "", "File containing game ids")
	err := comebacksCmd.MarkFlagRequired("idfile")
	if err != nil {
		panic(err)
	}
}
