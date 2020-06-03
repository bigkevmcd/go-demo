package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
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
			opts, err := redis.ParseURL(viper.GetString("redis_url"))
			if err != nil {
				return err
			}
			rdb := redis.NewClient(opts)

			h := demo.New(demo.Config{Redis: rdb, Key: viper.GetString("redis_key")})
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
		"redis_url",
		"redis://localhost:6379",
		"url to connect to Redis on",
	)
	logIfError(viper.BindPFlag("redis_url", cmd.Flags().Lookup("redis_url")))

	cmd.Flags().String(
		"redis_key",
		"demo:value",
		"key to fetch from Redis",
	)
	logIfError(viper.BindPFlag("redis_key", cmd.Flags().Lookup("redis_key")))
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
