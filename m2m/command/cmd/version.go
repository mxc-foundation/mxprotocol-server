package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the MXProtocal server version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
