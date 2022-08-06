/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/vaibhav135/go-cli-utils/data"
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	template := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: template,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word.",
		"What word would you like to make a note of?",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition.",
		fmt.Sprintf("What is the definition of %s?", word),
	}
	definition := promptGetInput(definitionPromptContent)
	categoryPromptContent := promptContent{
		"Please provide a category.",
		fmt.Sprintf("What category does %s belong to?", word),
	}
	category := promptGetSelect(categoryPromptContent)
	data.CreateNewNote(word, definition, category)
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		if index == -1 {
			items = append(items, result)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new studybuddy note",
	Long:  `Creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

func init() {
	noteCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
