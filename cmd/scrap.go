/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kirky/devy/internal/task"
	"kirky/devy/internal/utils"
	"log"

	"github.com/spf13/cobra"
)

// scrapCmd represents the scrap command
var scrapCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Automate clearing of temp dev files.",
	Long: `Run automation to clear generated temp dev files,
	which are not yet cleared.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("=====STARTING UP FILE CLEANUP=====")

		duration := "@every 10m"

		utils.RunJob(task.DeleteTempFiles, duration)
	},
}

func init() {
	rootCmd.AddCommand(scrapCmd)
}
