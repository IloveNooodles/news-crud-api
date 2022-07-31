package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func server() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
		},
	}
	return serverCmd
}
