/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// priorityCmd represents the priority command
var priorityCmd = &cobra.Command{
	Use:   "priority [ID] [low|medium|high]",
	Short: "设置任务优先级",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		newLevel := args[1]
		valid := map[string]bool{"low": true, "medium": true, "high": true}
		if !valid[newLevel] {
			fmt.Println("❗ 优先级无效，请使用 low、medium 或 high")
			return
		}
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Priority = newLevel
				saveTasks()
				fmt.Println("🔧 已设置任务优先级为:", newLevel)
				return
			}
		}
		fmt.Println("❗ 未找到指定任务")
	},
}

func init() {
	rootCmd.AddCommand(priorityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priorityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priorityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
