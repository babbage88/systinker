package cmd

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

// uptimeCmd represents the uptime command
var uptimeCmd = &cobra.Command{
	Use:   "uptime",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Getting system uptime...")
		currentUptime, _ := getUptime()
		fmt.Printf("Uptime: %s", currentUptime)
	},
}

func init() {
	rootCmd.AddCommand(uptimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uptimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uptimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getUptime() (time.Duration, error) {
	var info unix.Sysinfo_t
	if err := unix.Sysinfo(&info); err != nil {
		return time.Duration(0), err
	}
	return time.Duration(info.Uptime) * time.Second, nil
}
