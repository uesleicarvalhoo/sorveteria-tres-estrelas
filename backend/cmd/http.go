package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http"
)

func (app *application) httpServer() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "start a http server",
		Run: func(cmd *cobra.Command, args []string) {
			http.StartServer(app.config, app.db, app.kong)
		},
	}
}
