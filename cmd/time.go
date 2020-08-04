package cmd

import (
	"cobra/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}

			currentTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}

			currentTime, err = timer.GetCalculateTime(currentTime, duration)
			if err != nil {
				log.Fatalf("timer.GetCalculateTime err: %v", err)
			}
		}

		log.Printf("输入出结果： %s, %d \n", currentTime.Format(layout), currentTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(
		&calculateTime,
		"calculate",
		"c",
		"",
		"需要计算的时间，有效单位为时间戳或已格式化的时间",
	)

	calculateTimeCmd.Flags().StringVarP(
		&duration,
		"duration",
		"d",
		"",
		`持续时间，有效时间单位为： "ns", "us" (or "up"), "ms", "s", "m", "h"`,
	)
}
