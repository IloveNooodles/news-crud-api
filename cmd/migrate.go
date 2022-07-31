package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func migrate() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate into the sql",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("seeding called")
		},
	}
	return migrateCmd
}
