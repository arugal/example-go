package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	flags = struct {
		Addr     string
		ProjectA string
	}{
		ProjectA: "project-a",
	}

	rootCmd = &cobra.Command{
		Use:          "project-B",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectA := os.Getenv("PROJECTA_HOSTNAME")
			if projectA == "" {
				log.Printf("project-a use flages: %s", flags.ProjectA)
				projectA = flags.ProjectA
			}
			log.Printf("projct-a hostname:%s\n", projectA)

			router := gin.Default()

			router.GET("/v1", func(ctx *gin.Context) {
				resp, err := request(fmt.Sprintf("http://%s:%s/v1", projectA, "8081"))
				if err != nil {
					ctx.String(http.StatusInternalServerError, err.Error())
				} else {
					ctx.String(http.StatusOK, resp)
				}
			})
			return http.ListenAndServe(flags.Addr, router)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.Addr, "addr", ":8080", "project-B addr")
	rootCmd.PersistentFlags().StringVar(&flags.ProjectA, "projectA", "projectA", "project-A addr")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func request(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("request project-B error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
