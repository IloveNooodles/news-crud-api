package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func migrate() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "seeding",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("seeding called")
		},
	}
	return migrateCmd
}
