package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http"
)

var httpCommand = &cobra.Command{
	Use:   "http",
	Short: "start a http server",
	Run:   httpServerExecute,
}

func init() {
	rootCmd.AddCommand(httpCommand)
}

func httpServerExecute(cmd *cobra.Command, args []string) {
	http.StartServer()
}
