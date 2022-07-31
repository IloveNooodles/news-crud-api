package cmd

import (
	"github.com/spf13/cobra"
)

func Init() {
	rootCmd := &cobra.Command{
		Use:   "kumparan-techincal-test",
		Short: "kumparan technical test backend service",
	}

	rootCmd.AddCommand(
		server(),
		migrate(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
