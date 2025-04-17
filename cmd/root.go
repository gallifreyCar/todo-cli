/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Priority    string    `json:"priority"`
	RemindAt    time.Time `json:"remind_at"`
}

var tasks []Task

const taskFile = "tasks.json"

func loadTasks() {
	data, err := os.ReadFile(taskFile)
	if err == nil {
		_ = json.Unmarshal(data, &tasks)
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = os.WriteFile(taskFile, data, 0644)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "一个简单的任务管理 CLI 工具",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadTasks()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("运行失败:", err)
		os.Exit(1)
	}
	saveTasks()
}
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.awesomeProject7.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
