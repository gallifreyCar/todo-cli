/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [任务编号]",
	Short: "标记任务为完成",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❗ 任务编号无效")
			return
		}
		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				saveTasks()
				fmt.Println("🎉 已完成任务:", tasks[i].Description)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("❗ 未找到指定编号的任务")
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
