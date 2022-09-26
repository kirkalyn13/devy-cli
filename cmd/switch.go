/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kirky/devy/internal/task"
	"log"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch current git email logged in.",
	Long:  `Switch current git email logged in the device.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("=====SWITCHING GIT EMAIL=====")
		if len(args) != 1 {
			log.Print(noArgErr)
		} else {
			task.SwitchEmail(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
