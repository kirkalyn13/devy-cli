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

// unloadCmd represents the unload command
var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Clear docker cache.",
	Long:  `Clear docker cache from repeated development`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("=====STARTING UP DOCKER CLEANUP=====")

		duration := "@every 1h"

		utils.RunJob(task.UnloadDocker, duration)
	},
}

func init() {
	rootCmd.AddCommand(unloadCmd)
}
