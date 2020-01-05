package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	flags = struct {
		Addr string
	}{}

	rootCmd = &cobra.Command{
		Use:          "project-A",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			router := gin.Default()

			router.GET("/v1", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "hello, world")
			})
			return http.ListenAndServe(flags.Addr, router)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.Addr, "addr", ":8081", "project-A addr")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
