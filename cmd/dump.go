/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kirky/devy/internal/task"
	"log"

	"github.com/spf13/cobra"
)

const (
	noArgErr = "Invalid Arguments: No database specified."
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump dev schemas.",
	Long:  `Dump Schemas used in development.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("=====EXTRACTING DATABASE=====")
		if len(args) != 1 {
			log.Print(noArgErr)
		} else {
			if err := task.DumpDB(args[0]); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
