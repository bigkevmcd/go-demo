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

func makeRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-demo",
		Short: "Just a demo for GitOps",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts, err := redis.ParseURL(viper.GetString("redis_url"))
			cobra.CheckErr(err)

			rdb := redis.NewClient(opts)

			h := demo.New(demo.Config{Redis: rdb, Key: viper.GetString("redis_key")})

			addr := fmt.Sprintf(":%d", viper.GetInt("port"))
			log.Printf("listening on %s, connecting to Redis %s\n", addr, opts.Addr)

			return http.ListenAndServe(addr, h)
		},
	}
	cmd.Flags().Int(
		"port",
		8080,
		"port to serve requests on",
	)
	cobra.CheckErr(viper.BindPFlag("port", cmd.Flags().Lookup("port")))

	cmd.Flags().String(
		"redis_url",
		"redis://localhost:6379",
		"url to connect to Redis on",
	)
	cobra.CheckErr(viper.BindPFlag("redis_url", cmd.Flags().Lookup("redis_url")))

	cmd.Flags().String(
		"redis_key",
		"demo:value",
		"key to fetch from Redis",
	)
	cobra.CheckErr(viper.BindPFlag("redis_key", cmd.Flags().Lookup("redis_key")))

	return cmd
}

func initConfig() {
	viper.AutomaticEnv()
}

// Execute is the main entry point into this component.
func Execute() {
	cobra.CheckErr(makeRootCmd().Execute())
}
