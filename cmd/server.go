package cmd

import (
	"github.com/IloveNooodles/kumparan-techincal-test/config"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/routes"
	"github.com/IloveNooodles/kumparan-techincal-test/pkg/lib"
	svr "github.com/IloveNooodles/kumparan-techincal-test/pkg/server"
	"github.com/spf13/cobra"
)

func server() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "run the server",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Init(config.C.Db_url)
			service := svr.NewServer()
			routes.RoutesInit(service.App(), lib.DB)
			service.Start(config.C.Port)
		},
	}
	return serverCmd
}
