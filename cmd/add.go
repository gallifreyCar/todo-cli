/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nextID int

func getNextID() int {
	maxN := 0
	for _, t := range tasks {
		if t.ID > maxN {
			maxN = t.ID
		}
	}
	return maxN + 1
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [任务内容]",
	Short: "添加一个新任务",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := Task{
			ID:          getNextID(),
			Description: args[0],
			Priority:    "medium",
		}
		tasks = append(tasks, task)
		fmt.Println("✅ 添加任务:", task.Description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
