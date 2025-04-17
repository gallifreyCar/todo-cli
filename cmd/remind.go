/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// remindCmd represents the remind command
var remindCmd = &cobra.Command{
	Use:   "remind",
	Short: "启动本地提醒调度器，定时检查并提醒任务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⏰ 提醒调度启动，每 10 秒扫描一次任务...")

		for {
			now := time.Now()
			for _, task := range tasks {
				// 满足提醒条件：未完成、有设置提醒时间、且当前时间已超过提醒时间但不超过1分钟（避免多次提醒）
				if !task.Done && !task.RemindAt.IsZero() {
					diff := now.Sub(task.RemindAt)
					if diff >= 0 && diff <= time.Minute {
						fmt.Printf("🔔 [提醒] [%s] %s\n", task.RemindAt.Format("15:04"), task.Description)
					}
				}
			}
			time.Sleep(10 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(remindCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
