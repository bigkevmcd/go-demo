package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bigkevmcd/go-demo/pkg/demo"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func logIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func makeRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-demo",
		Short: "Just a demo for GitOps",
		RunE: func(cmd *cobra.Command, args []string) error {
			h := demo.New(demo.Config{Name: viper.GetString("name")})
			addr := fmt.Sprintf(":%d", viper.GetInt("port"))
			log.Printf("listening on %s\n", addr)
			return http.ListenAndServe(addr, h)
		},
	}
	cmd.Flags().Int(
		"port",
		8080,
		"port to serve requests on",
	)
	logIfError(viper.BindPFlag("port", cmd.Flags().Lookup("port")))

	cmd.Flags().String(
		"name",
		"default",
		"name to serve responses with",
	)
	logIfError(viper.BindPFlag("name", cmd.Flags().Lookup("name")))
	return cmd
}

func initConfig() {
	viper.AutomaticEnv()
}

// Execute is the main entry point into this component.
func Execute() {
	if err := makeRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
