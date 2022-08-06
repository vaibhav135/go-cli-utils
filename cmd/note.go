/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note can be anything you'd like to study and review.",
	Long:  `A note can be anything you'd like to study and review.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
