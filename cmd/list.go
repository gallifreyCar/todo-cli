/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有任务",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) == 0 {
			fmt.Println("📭 没有任务")
			return
		}
		for _, task := range tasks {
			status := "❌"
			if task.Done {
				status = "✅"
			}
			remind := ""
			if !task.RemindAt.IsZero() {
				remind = task.RemindAt.Format("2006-01-02 15:04")
			}
			fmt.Printf("[%d] %s [%s] %s (提醒: %s)\n", task.ID, status, task.Priority, task.Description, remind)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
